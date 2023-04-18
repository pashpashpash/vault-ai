// @flow
import * as React from 'react';
import s from './index.less';

import Go from '../Go';
import URL from 'url-parse';

const Header = (): React.Node => {
    const curHost = window.location.hostname;
    let thisSite = curHost;

    if (curHost === 'localhost') {
        thisSite = 'http://localhost:8100';
    } else {
        thisSite = 'https://34.82.254.177/';
    }

    const ReedemableLinks = [
        {
            to: thisSite + '/',
            text: 'Home',
        },
    ];

    const [showPancakeMenu, setShowPancakeMenu] = React.useState(false);

    const links = ReedemableLinks.map(
        (linkData: { to: string, text: string }, i: number): React.Node => (
            <Go
                to={linkData.to}
                key={i}
                className={s.menuButton}
                data-category={'Giftcard-header'}
                data-action={linkData.text}>
                {linkData.text}
            </Go>
        )
    );

    const submenuClass = [s.sublinks];
    if (showPancakeMenu) {
        submenuClass.push(s.visible);
    }

    const HamburgerMenu = (
        <div className={s.hamburgerMenu}>
            <div
                className={s.pancakeIcon}
                onClick={() => {
                    setShowPancakeMenu(!showPancakeMenu);
                }}
            />

            <div className={submenuClass.join(' ')}>{links}</div>
        </div>
    );

    return (
        <div className={s.header}>
            {HamburgerMenu}
            <Go
                to={thisSite}
                data-action={'home'}
                data-category={'Giftcard-header'}>
                <div className={s.giftLogo}>The Vault</div>
            </Go>
        </div>
    );
};

export default Header;
