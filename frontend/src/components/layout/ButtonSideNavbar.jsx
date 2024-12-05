import styled from "styled-components";

// eslint-disable-next-line react/prop-types
function ButtonSideNavbar({ Icon, active }) {
    return (
        <ButtonContainer>
            {active && <ActiveBar />}
            <IconStyled active={active} as={Icon} />
        </ButtonContainer>
    );
}

const ButtonContainer = styled.div`
    width: 60px;
    height: 50px;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 10px;
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
    background-color: var(--color-fennect);
    left: 0;
`;

export default ButtonSideNavbar;