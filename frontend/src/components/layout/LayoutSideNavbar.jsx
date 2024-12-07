/* eslint-disable react/prop-types */
import styled from "styled-components";
import { Link, useLocation } from "react-router-dom";
import { RiHome2Fill } from "react-icons/ri";
import { MdAccountCircle, MdSettings, MdGroup, MdNotifications } from "react-icons/md";
import { RiSearchFill } from "react-icons/ri";
import { BiSolidMessageRounded } from "react-icons/bi";
import ButtonSideNavbar from "./ButtonSideNavbar";

const LayoutSideNavbar = ({ scrollDirection }) => {
    const location = useLocation();

    return (
        <SideNavbarContainer style={{ bottom: scrollDirection === "up" ? "-60px" : "0" }}>
            <StyledLink to="/" className="home-link">
                <ButtonSideNavbar
                    Icon={RiHome2Fill}
                    active={location.pathname === "/"}
                    label="Home"
                    badgeNum={0}
                />
            </StyledLink>
            <StyledLink to="/search">
                <ButtonSideNavbar
                    Icon={RiSearchFill}
                    active={location.pathname === "/search"}
                    label="Search"
                />
            </StyledLink>
            <StyledLink to="/messages">
                <ButtonSideNavbar
                    Icon={BiSolidMessageRounded}
                    active={location.pathname === "/messages"}
                    badgeNum={2}
                />
            </StyledLink>
            <StyledLink to="/notifications">
                <ButtonSideNavbar
                    Icon={MdNotifications}
                    active={location.pathname === "/notifications"}
                    badgeNum={9}
                />
            </StyledLink>
            <StyledLink to="/groups">
                <ButtonSideNavbar
                    Icon={MdGroup}
                    active={location.pathname === "/groups"}
                />
            </StyledLink>

            <NonMobileContainer>
                <StyledLink to="/account">
                    <ButtonSideNavbar
                        Icon={MdAccountCircle}
                        active={location.pathname === "/account"}
                    />
                </StyledLink>
                <StyledLink to="/settings">
                    <ButtonSideNavbar
                        Icon={MdSettings}
                        active={location.pathname === "/settings"}
                    />
                </StyledLink>
            </NonMobileContainer>
        </SideNavbarContainer>
    );
};

const SideNavbarContainer = styled.div`
    background-color: var(--color-secondary);
    height: 100%;
    width: 60px;
    min-width: 60px;
    display: flex;
    flex-direction: column;
    justify-content: start;
    align-items: center;
    transition: bottom 0.5s;

    @media (max-width: 600px) {
        height: 50px;
        width: 100%;
        flex-direction: row;
        justify-content: center;
        position: fixed;
        bottom: 0;
    }
`;

const NonMobileContainer = styled.div`
    display: flex;
    flex-direction: column;
    margin-top: auto;

    @media (max-width: 600px) {
        display: none;
    }
`;

const StyledLink = styled(Link)`
    @media (max-width: 600px) {
        margin-inline: auto;
    }

    &.home-link {
        @media (max-width: 600px) {
            order: 3;
        }
    }

    &:nth-child(2) {
        @media (max-width: 600px) {
            order: 4;
        }
    }

    &:nth-child(3) {
        @media (max-width: 600px) {
            order: 5;
        }
    }

    &:nth-child(4) {
        @media (max-width: 600px) {
            order: 2;
        }
    }

    &:nth-child(5) {
        @media (max-width: 600px) {
            order: 1;
        }
    }
`;

export default LayoutSideNavbar;