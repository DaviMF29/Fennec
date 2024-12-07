/* eslint-disable react/prop-types */
import styled from "styled-components";
import IconComponent from '../IconComponent'
import { MdMenu } from "react-icons/md";

const LayoutTopBarContainer = styled.div`
    display: flex;
    width: 100%;
    height: 30px;
    background-color: var(--color-secondary);
    transition: top 0.5s;

    @media (max-width: 600px) {
        height: 50px;
        position: fixed;
        background-color: var(--color-primary);
    }
`;

const IconMenu = styled(MdMenu)`
    font-size: 2em;
    display: none;
    cursor: pointer;
    width: 60px;
    margin-block: auto;
    opacity: 0.5;
    transition: all 0.3s;

    &:hover{
        opacity: 1;
    }

    @media (max-width: 600px) {
        display: flex;
    }
`

const IconContainer = styled.div`
    display: flex;
    width: 60px;
    height: 100%;
    justify-content: center;
    align-items: center;
    opacity: 0.5;
    transition: all 0.3s;

    &:hover{
        opacity: 1;
    }

    @media (max-width: 600px) {
        margin-inline: auto;
        height: 80%;
        margin-block: auto;
    }
`

const MobileSpacer = styled.div`
    width: 60px;
`

function LayoutTopBar({ scrollDirection }) {
    return (
        <LayoutTopBarContainer style={{ top: scrollDirection === "up" ? "-50px" : "0" }}>
            <IconMenu />
            <IconContainer>
                <IconComponent height={'70%'} fill={'var(--color-text-primary)'} />
            </IconContainer>
            <MobileSpacer />
        </LayoutTopBarContainer>
    );
}

export default LayoutTopBar;