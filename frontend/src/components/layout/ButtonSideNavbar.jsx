/* eslint-disable react/prop-types */
import styled from "styled-components";

function ButtonSideNavbar({ Icon, active, badgeNum }) {
    return (
        <ButtonContainer>
            {active && <ActiveBar />}
            <IconStyled active={active ? "true" : "false"} as={Icon} />
            {badgeNum >= 1 && badgeNum < 10 && (
                <Badge>{badgeNum}</Badge>
            )}
            {badgeNum === 0 || badgeNum > 9 && (
                <BadgeEmpty />
            )}
        </ButtonContainer>
    );
}

const ButtonContainer = styled.div`
    width: 60px;
    height: 50px;
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
    transition: all 0.3s;

    &::before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        width: 0;
        height: 0;
        background-color: transparent;
        opacity: 0;
        width: 60px;
        height: 50px;
        transition: all 0.3s;
    }

    &:hover {
        &::before {
            background-color: var(--color-text-primary);
            opacity: 0.1;
        }
    }

    @media (max-width: 600px) {
        width: auto;

        &:hover {
            &::before {
                display: none;
            }
        }
    }
`;

const IconStyled = styled.div`
    font-size: 1.6em;
    transition: all 0.3s;
    opacity: ${props => (props.active === "true" ? 1 : 0.45)};

    @media (max-width: 600px) {
        margin-inline: auto;
    }

    @media (max-width: 400px) {
        font-size: 1.3em;
    }

    @media (max-width: 200px) {
        font-size: 1em;
    }
`;

const ActiveBar = styled.div`
    position: absolute;
    width: 2px;
    height: 50px;
    background-color: var(--color-accent);
    left: 0;

    @media (max-width: 600px) {
        display: none;
    }
`;

const Badge = styled.div`
    position: absolute;
    right: 10px;
    bottom: 25px;
    width: 17px;
    height: 17px;
    border-radius: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 0.7em;
    font-weight: bold;
    padding: auto;
    background-color: var(--color-accent);
    z-index: 10;
    user-select: none;
`;

const BadgeEmpty = styled.div`
    position: absolute;
    right: 15px;
    bottom: 35px;
    width: 7px;
    height: 7px;
    border-radius: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 0.7em;
    font-weight: bold;
    padding: auto;
    background-color: var(--color-accent);
    z-index: 10;
`;

export default ButtonSideNavbar;