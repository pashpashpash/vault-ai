// @flow
import * as React from 'react';
import { Link } from 'react-router-dom';
import URL from 'url-parse';

type Props = {|
    className?: string,
    to: ?string,
    external?: ?boolean,
    'data-action': string,
    'data-category': string,
    children?: ?React.Node,
    onClick?: () => void,
|};
let MY_HOSTNAMES = ['localhost', 'vault.app'];

/**
 * <Go to="url" />
 * Combines <Link> and <a> into a simple hassle-free interface
 */
class Go extends React.Component<Props> {
    render(): React.Node {
        const curHost = window.location.hostname;
        const toUrlParse = new URL(this.props.to);

        if (this.props.to == null) {
            return (
                <a target="_blank" rel="noopener noreferrer" {...this.props} />
            );
        }

        if (
            MY_HOSTNAMES.indexOf(toUrlParse.hostname) >= 0 &&
            !this.props.external
        ) {
            const newProps = Object.assign({}, this.props, {
                to: toUrlParse.pathname + toUrlParse.hash + toUrlParse.query,
            });
            return <Link {...newProps} />;
        } else {
            return (
                <a
                    href={this.props.to}
                    target="_blank"
                    rel="noopener noreferrer"
                    {...this.props}
                />
            );
        }
    }
}

export default Go;
