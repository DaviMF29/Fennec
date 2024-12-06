import styled from "styled-components";
import { Link, useLocation } from "react-router-dom";
import { RiHome2Fill } from "react-icons/ri";
import { MdAccountCircle, MdSettings, MdSearch, MdGroup, MdNotifications } from "react-icons/md";
import { BiSolidMessageRounded } from "react-icons/bi";
import ButtonSideNavbar from "./ButtonSideNavbar";

const LayoutSideNavbar = () => {
    const location = useLocation();

    return (
        <SideNavbarContainer>
            <Link to="/">
                <ButtonSideNavbar
                    Icon={RiHome2Fill}
                    active={location.pathname === "/"}
                    label="Home"
                    badgeNum={0}
                />
            </Link>
            <Link to="/search">
                <ButtonSideNavbar
                    Icon={MdSearch }
                    active={location.pathname === "/search"}
                    label="Search"
                />
            </Link>
            <Link to="/messages">
                <ButtonSideNavbar
                    Icon={BiSolidMessageRounded}
                    active={location.pathname === "/messages"}
                    badgeNum={2}
                />
            </Link>
            <Link to="/notifications">
                <ButtonSideNavbar
                    Icon={MdNotifications}
                    active={location.pathname === "/notifications"}
                    badgeNum={5}
                />
            </Link>
            <Link to="/groups">
                <ButtonSideNavbar
                    Icon={MdGroup}
                    active={location.pathname === "/groups"}
                />
            </Link>

            <Spacer />

            <Link to="/account">
                <ButtonSideNavbar
                    Icon={MdAccountCircle}
                    active={location.pathname === "/account"}
                />
            </Link>
            <Link to="/settings">
                <ButtonSideNavbar
                    Icon={MdSettings}
                    active={location.pathname === "/settings"}
                />
            </Link>
        </SideNavbarContainer>
    );
};

const Spacer = styled.div`
    height: 100%;
    width: 0;
`;

const SideNavbarContainer = styled.div`
    background-color: var(--color-secondary);
    height: 100%;
    width: 60px;
    min-width: 60px;
    display: flex;
    flex-direction: column;
    justify-content: start;
    align-items: center;
`;

export default LayoutSideNavbar;