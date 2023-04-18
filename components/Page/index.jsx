// @flow
import * as React from 'react';
import PropTypes from 'prop-types';
// import Footer from '../Footer';
import s from './index.less';
import '../common.less';
import Footer from '../Footer';
type Props = {|
    contentPreferredWidth: number,
    anchorTarget?: string,
    pageClass?: string,
    pageBackgroundClass?: string,
    contentClass?: string,
    contentBackgroundClass?: string,
    contentBackground?: string,
    footerComponent?: React.Node,
    showContentFooter?: boolean,
    verticalCenter?: boolean,
    footerEnabled?: boolean,
    footerFloating?: boolean,
    heightOverride?: number,
    noTopPadding?: boolean,
    style?: Object,
    children?: React.Node,
|};

const Page = (props: Props): React.Node => {
    const pageClass = props.pageClass != null ? props.pageClass : '';
    const pageBackgroundClass = props.pageBackgroundClass != null ? props.pageBackgroundClass : '';
    const contentClass = props.contentClass != null ? props.contentClass : '';
    const contentBackgroundClass =
        props.contentBackgroundClass != null
            ? props.contentBackgroundClass
            : '';

    const footerClass = props.showContentFooter === true ? s.showFooter : '';
    const centerClass = props.verticalCenter === true ? s.verticalCenter : '';
    const hasFooter = props.footerEnabled === true;
    const floatingFooter = props.footerFloating === true;
    const backgroundStyles = {
        backgroundImage: props.contentBackground
            ? 'url(' + props.contentBackground + ')'
            : null,
    };
    let pageStyles = {};
    const contentStyles: { width: number, paddingTop?: string | number } = {
        width: props.contentPreferredWidth,
    };

    if (props.heightOverride != null) {
        pageStyles.height = props.heightOverride + 'px';
        pageStyles.minHeight = props.heightOverride + 'px';
    }

    if (props.noTopPadding === true) {
        contentStyles.paddingTop = '0px';
    }
    if (props.style) {
        pageStyles = {
            ...pageStyles,
            ...props.style,
        };
    }

    return (
        <div>
            <div
                id={props.anchorTarget}
                style={pageStyles}
                className={[s.page, pageClass, pageBackgroundClass].join(' ')}>
                <div
                    className={[
                        s.contentBackground,
                        contentBackgroundClass,
                    ].join(' ')}
                    style={backgroundStyles}>
                    <div
                        className={[
                            s.content,
                            contentClass,
                            footerClass,
                            centerClass,
                        ].join(' ')}
                        style={contentStyles}>
                        {props.children}
                    </div>
                </div>
            </div>
            {hasFooter &&
                    (props.footerComponent != null ? (
                        props.footerComponent
                    ) : (
                        <Footer floating={floatingFooter} backgroundClass={pageBackgroundClass}/>
                    ))}
        </div>
    );
};

Page.propTypes = {
    contentPreferredWidth: PropTypes.number.isRequired,
    showContentFooter: PropTypes.bool,
    pageClass: PropTypes.string,
    contentClass: PropTypes.string,
    contentBackgroundClass: PropTypes.string,
    contentBackground: PropTypes.string,
    verticalCenter: PropTypes.bool,
    footerEnabled: PropTypes.bool,
    footerFloating: PropTypes.bool,
    anchorTarget: PropTypes.string,
    heightOverride: PropTypes.number,
    noTopPadding: PropTypes.bool,
    footerComponent: PropTypes.node,
};

export default Page;
