import styled from "styled-components";
import IconComponent from '../IconComponent'

const LayoutTopBarContainer = styled.div`
    display: flex;
    width: 100%;
    height: 30px;
    background-color: var(--color-secondary);
`;

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
`

function LayoutTopBar(){
    return(
        <LayoutTopBarContainer>
            <IconContainer>
                <IconComponent height={'70%'} width={'auto'} fill={'var(--color-text-primary)'}/>
            </IconContainer>
        </LayoutTopBarContainer>
    )
}

export default LayoutTopBar;