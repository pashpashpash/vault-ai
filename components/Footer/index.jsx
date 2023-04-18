// @flow
import * as React from 'react';
import s from './index.less';
import PropTypes from 'prop-types';
import Go from '../Go';

type Props = {
    backgroundOnly?: boolean,
    backgroundClass?: string,
};

const Footer = (props: Props): React.Node => {
    const footerClasses = [s.footer];
    if (props.backgroundClass) footerClasses.push(props.backgroundClass);
    if (props.backgroundOnly) footerClasses.push(s.backgroundOnly);
    return (
        <div className={footerClasses.join(' ')}>
            <div className={s.footerContainer}>
                <div className={[s.cropHeight, s.flip].join(' ')}></div>
                {props.backgroundOnly === true ? null : (
                    <div className={s.content}>
                        <div className={s.linkRow}>
                            <div className={s.socialButtons}></div>
                            <div className={s.linkColumn}>
                                <div className={s.linkHeader}>Vault</div>
                                <Go
                                    to={'/'}
                                    data-action={'vault-home'}
                                    data-category={'footer-link'}>
                                    Home
                                </Go>
                            </div>
                            <div className={s.linkColumn}>
                                <div className={s.linkHeader}>Help</div>
                                <div>FAQ</div>
                                <div>Contact Us</div>
                            </div>
                            <div className={s.linkColumn}></div>
                            <div className={s.linkColumn}></div>
                        </div>
                        <div className={s.bottomContent}>
                            <img
                                style={{ display: 'inline-block' }}
                                height="40px"
                                width="40px"
                                src="/img/logos/vault-favicon.png"
                            />

                            <span>
                                © pashpashpash {new Date().getFullYear()} All
                                Rights Reserved.
                            </span>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
};
Footer.propTypes = {
    backgroundOnly: PropTypes.bool,
};

export default Footer;
