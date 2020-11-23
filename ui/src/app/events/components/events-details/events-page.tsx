import {Page, SlidingPanel, Tabs} from 'argo-ui';
import {useContext, useEffect, useState} from 'react';
import React = require('react');
import {RouteComponentProps} from 'react-router-dom';
import {Observable} from 'rxjs';
import {kubernetes, Workflow} from '../../../../models';
import {EventSource} from '../../../../models/event-source';
import {Sensor} from '../../../../models/sensor';
import {uiUrl} from '../../../shared/base';
import {ErrorNotice} from '../../../shared/components/error-notice';
import {GraphPanel} from '../../../shared/components/graph/graph-panel';
import {Node} from '../../../shared/components/graph/types';
import {NamespaceFilter} from '../../../shared/components/namespace-filter';
import {ResourceEditor} from '../../../shared/components/resource-editor/resource-editor';
import {ZeroState} from '../../../shared/components/zero-state';
import {Context} from '../../../shared/context';
import {historyUrl} from '../../../shared/history';
import {ListWatch} from '../../../shared/list-watch';
import {services} from '../../../shared/services';
import {EventsPanel} from '../../../workflows/components/events-panel';
import {FullHeightLogsViewer} from '../../../workflows/components/workflow-logs-viewer/full-height-logs-viewer';
import {buildGraph} from './build-graph';
import {genres} from './genres';
import {ID} from './id';

require('./event-page.scss');

export const EventsPage = (props: RouteComponentProps<any>) => {
    // boiler-plate
    const {navigation} = useContext(Context);
    const {match, location, history} = props;
    const queryParams = new URLSearchParams(location.search);

    // state for URL and query parameters
    const [namespace, setNamespace] = useState(match.params.namespace || '');
    const [showFlow, setShowFlow] = useState(queryParams.get('showFlow') === 'true');
    const [showWorkflows, setShowWorkflows] = useState(queryParams.get('showWorkflows') === 'true');
    const [expanded, setExpanded] = useState(queryParams.get('expanded') === 'true');
    const [selectedNode, setSelectedNode] = useState<Node>(queryParams.get('selectedNode'));
    const [tab, setTab] = useState<Node>(queryParams.get('tab'));
    useEffect(
        () =>
            history.push(
                historyUrl('events/{namespace}', {
                    namespace,
                    showFlow,
                    showWorkflows,
                    expanded,
                    selectedNode,
                    tab
                })
            ),
        [namespace, showFlow, showWorkflows, expanded, expanded, tab]
    );

    // internal state
    const [error, setError] = useState<Error>();
    const [eventSources, setEventSources] = useState<EventSource[]>();
    const [sensors, setSensors] = useState<Sensor[]>();
    const [workflows, setWorkflows] = useState<Workflow[]>();
    const [flow, setFlow] = useState<{[id: string]: {count: number; timeout?: any}}>({}); // event flowing?

    // when namespace changes, we must reload
    useEffect(() => {
        const listWatch = new ListWatch<EventSource>(
            () => services.eventSource.list(namespace),
            resourceVersion => services.eventSource.watch(namespace, resourceVersion),
            () => setError(null),
            () => setError(null),
            items => setEventSources([...items]),
            setError
        );
        listWatch.start();
        return () => listWatch.stop();
    }, [namespace]);
    useEffect(() => {
        const listWatch = new ListWatch<Sensor>(
            () => services.sensor.list(namespace),
            resourceVersion => services.sensor.watch(namespace, resourceVersion),
            () => setError(null),
            () => setError(null),
            items => setSensors([...items]),
            setError
        );
        listWatch.start();
        return () => listWatch.stop();
    }, [namespace]);
    useEffect(() => {
        if (!showWorkflows) {
            setWorkflows(null);
            return;
        }
        const listWatch = new ListWatch<Workflow>(
            () =>
                services.workflows.list(namespace, null, ['events.argoproj.io/sensor', 'events.argoproj.io/trigger'], null, [
                    'metadata',
                    'items.metadata.name',
                    'items.metadata.namespace',
                    'items.metadata.creationTimestamp',
                    'items.metadata.labels'
                ]),
            resourceVersion =>
                services.workflows.watch({
                    namespace,
                    resourceVersion,
                    labels: ['events.argoproj.io/sensor', 'events.argoproj.io/trigger']
                }),
            () => setError(null),
            () => setError(null),
            (items, item, type) => {
                setWorkflows([...items]);
                if (type === 'ADDED') {
                    markFlowing(ID.join('Workflow', item.metadata.namespace, item.metadata.name));
                }
            },
            setError
        );
        listWatch.start();
        return () => listWatch.stop();
    }, [namespace, showWorkflows]);
    // follow logs and mark flow
    const markFlowing = (id: Node) => {
        if (!flow) {
            return;
        }
        setError(null);
        setFlow(newFlow => {
            if (!newFlow[id]) {
                newFlow[id] = {count: 0};
            }
            clearTimeout(newFlow[id].timeout);
            newFlow[id].count++;
            newFlow[id].timeout = setTimeout(() => {
                setFlow(evenNewerFlow => {
                    delete evenNewerFlow[id].timeout;
                    return Object.assign({}, evenNewerFlow); // Object.assign work-around to make sure state updates
                });
            }, 3000);
            return Object.assign({}, newFlow);
        });
    };
    useEffect(() => {
        if (!showFlow) {
            return;
        }
        const sub = services.eventSource
            .eventSourcesLogs(namespace, '', '', '', 'dispatching.*event', 0)
            .filter(e => !!e && !!e.eventSourceName)
            .subscribe(e => markFlowing(ID.join('EventSource', e.namespace, e.eventSourceName, e.eventName)), setError);
        return () => sub.unsubscribe();
    }, [namespace, showFlow]);
    useEffect(() => {
        if (!showFlow) {
            return;
        }
        const sub = services.sensor
            .sensorsLogs(namespace, '', '', 'successfully processed', 0)
            .filter(e => !!e)
            .subscribe(e => {
                markFlowing(ID.join('Sensor', e.namespace, e.sensorName));
                if (e.triggerName) {
                    markFlowing(ID.join('Trigger', e.namespace, e.sensorName, e.triggerName));
                }
            }, setError);
        return () => sub.unsubscribe();
    }, [namespace, showFlow]);

    const graph = buildGraph(eventSources, sensors, workflows, flow, expanded);

    const selected = (() => {
        if (!selectedNode) {
            return;
        }
        const x = ID.split(selectedNode);
        const kind = x.type === 'EventSource' ? 'EventSource' : 'Sensor';
        const resources: {metadata: kubernetes.ObjectMeta}[] = (kind === 'EventSource' ? eventSources : sensors) || [];
        const value = resources.find((y: {metadata: kubernetes.ObjectMeta}) => y.metadata.namespace === x.namespace && y.metadata.name === x.name);
        return {kind, value, ...x};
    })();

    return (
        <Page
            title='Events'
            toolbar={{
                actionMenu: {
                    items: [
                        {
                            action: () => setShowFlow(!showFlow),
                            iconClassName: showFlow ? 'fa fa-toggle-on' : 'fa fa-toggle-off',
                            title: 'Show event-flow'
                        },
                        {
                            action: () => setShowWorkflows(!showWorkflows),
                            iconClassName: showWorkflows ? 'fa fa-toggle-on' : 'fa fa-toggle-off',
                            title: 'Show workflows'
                        },
                        {
                            action: () => setExpanded(!expanded),
                            iconClassName: expanded ? 'fa fa-compress' : 'fa fa-expand',
                            title: 'Collapse/expand hidden nodes'
                        }
                    ]
                },
                tools: [<NamespaceFilter key='namespace-filter' value={namespace} onChange={setNamespace} />]
            }}>
            <ErrorNotice error={error} style={{margin: 20}} />
            {graph.nodes.size === 0 ? (
                <ZeroState>
                    <p>Argo Events allow you to trigger workflows, lambadas, and other actions when an event such as a webhooks, message, or a cron schedule occurs.</p>
                    <p>
                        <a href='https://argoproj.github.io/argo-events/'>Learn more</a>
                    </p>
                </ZeroState>
            ) : (
                <>
                    <GraphPanel
                        storageScope='events'
                        classNames='events'
                        graph={graph}
                        nodeGenres={genres}
                        nodeClassNames={{'': true, 'Pending': true, 'Ready': true, 'Running': true, 'Failed': true, 'Succeeded': true, 'Error': true}}
                        iconShapes={{workflow: 'circle', collapsed: 'circle', conditions: 'circle'}}
                        horizontal={true}
                        selectedNode={selectedNode}
                        onNodeSelect={x => {
                            const id = ID.split(x);
                            if (id.type === 'Workflow') {
                                navigation.goto(uiUrl('workflows/' + id.namespace + '/' + id.name));
                            } else if (id.type === 'Collapsed') {
                                setExpanded(true);
                            } else {
                                setSelectedNode(x);
                            }
                        }}
                    />
                    {showFlow && (
                        <p className='argo-container'>
                            <i className='fa fa-info-circle' /> Event-flow is proxy for events. It is based on the pod logs of the event sources and sensors, so should be treated
                            only as indicative of activity.
                        </p>
                    )}
                </>
            )}
            <SlidingPanel isShown={!!selectedNode} onClose={() => setSelectedNode(null)}>
                {!!selectedNode && (
                    <div>
                        <h4>
                            {selected.kind}/{selected.name}
                        </h4>
                        <h5>{selected.key}</h5>
                        <Tabs
                            navTransparent={true}
                            selectedTabKey={tab}
                            onTabSelected={setTab}
                            tabs={[
                                {
                                    title: 'SUMMARY',
                                    key: 'summary',
                                    content: <ResourceEditor kind={selected.kind} value={selected.value} />
                                },
                                {
                                    title: 'LOGS',
                                    key: 'logs',
                                    content: (
                                        <div className='white-box' style={{height: 600}}>
                                            <FullHeightLogsViewer
                                                source={{
                                                    key: 'logs',
                                                    loadLogs: () =>
                                                        ((selected.kind === 'Sensor'
                                                            ? services.sensor.sensorsLogs(namespace, selected.name, selected.key, '', 50)
                                                            : services.eventSource.eventSourcesLogs(namespace, selected.name, '', selected.key, '', 50)) as Observable<any>)
                                                            .filter(e => !!e)
                                                            .map(
                                                                e =>
                                                                    Object.entries(e)
                                                                        .map(([key, value]) => key + '=' + value)
                                                                        .join(', ') + '\n'
                                                            ),
                                                    shouldRepeat: () => false
                                                }}
                                            />
                                        </div>
                                    )
                                },
                                {
                                    title: 'EVENTS',
                                    key: 'events',
                                    content: <EventsPanel kind={selected.kind} namespace={selected.namespace} name={selected.name} />
                                }
                            ]}
                        />
                    </div>
                )}
            </SlidingPanel>
        </Page>
    );
};
