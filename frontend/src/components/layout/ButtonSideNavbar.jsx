import styled from "styled-components";

// eslint-disable-next-line react/prop-types
function ButtonSideNavbar({ Icon, active, badgeNum }) {
    return (
        <ButtonContainer>
            {active && <ActiveBar/>}
            <IconStyled active={active} as={Icon} />
            {badgeNum >= 1 && (
                <Badge>
                    {badgeNum}
                </Badge>
            )}
            {badgeNum == 0 && (
                <BadgeEmpty/>
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
`;

const IconStyled = styled.div`
    font-size: 1.6em;
    transition: all 0.3s;
    opacity: ${props => (props.active ? 1 : 0.7)};
`;

const ActiveBar = styled.div`
    position: absolute;
    width: 2px;
    height: 50px;
    background-color: var(--color-accent);
    left: 0;
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
`

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
`

export default ButtonSideNavbar;