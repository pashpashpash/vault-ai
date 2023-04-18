/* eslint-disable flowtype/require-return-type */
// @flow
import * as React from 'react';
import { lazy, Suspense } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

const LandingPage = lazy(() => import('./Pages/LandingPage'));
const Header = lazy(() => import('./Header/index'));

const loading = (
    <div
        style={{
            color: '#fff',
            minHeight: '100vh',
            backgroundColor: 'white',
            minWidth: '100%',
        }}>
        loading...
    </div>
);

import s from './routes.less';

type Props = {|
    change?: Object,
|};

type State = {|
    error: Object,
|};

class NugRoute extends Route {
    unifiedHandler(props: Props) {
        if (typeof props.change === 'function') {
            props.change(props);
        }
    }

    componentDidMount() {
        this.unifiedHandler(this.props);
    }

    componentWillReceiveProps() {
        this.unifiedHandler(this.props);
    }
}

class AppRouting extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            error: null,
        };

        this.handlePageChange = this.handlePageChange.bind(this);
    }

    componentDidCatch(error: Error, info: Object) {
        this.setState({
            error: {
                message: error.message,
                errorStack: error.stack,
                reactStack: info.componentStack,
            },
        });
    }

    handlePageChange: () => null = () => {
        window.scrollTo(0, 0);
        return null;
    };

    render(): React.Node {
        if (this.state.error != null) {
            return (
                <div className={s.crashPage}>
                    <div className={s.errorBox}>
                        <div className={s.title}>
                            Oops, A Wild Error Appeared! It says:
                        </div>

                        <div className={s.message}>
                            {'"' + this.state.error.message + '"'}
                        </div>

                        <div className={s.reactStack}>
                            {this.state.error.reactStack}
                        </div>

                        <div className={s.errorStack}>
                            {this.state.error.errorStack}
                        </div>
                    </div>
                </div>
            );
        } else {
            return (
                <Suspense fallback={loading}>
                    <BrowserRouter>
                        <div>
                            <Switch>
                                <NugRoute component={Header} />
                            </Switch>
                            <Switch>
                                <NugRoute
                                    exact
                                    path="/"
                                    component={LandingPage}
                                    change={this.handlePageChange}
                                />
                            </Switch>
                        </div>
                    </BrowserRouter>
                </Suspense>
            );
        }
    }
}

export default AppRouting;
