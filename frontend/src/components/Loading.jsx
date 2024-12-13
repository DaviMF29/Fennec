import styled, { keyframes } from 'styled-components';
import Iconcomponent from './IconComponent'

const Container = styled.div`
    display: flex;
    width: 100vw;
    height: 100vh;
    align-items: center;
    justify-content: center;
    overflow: hidden;
`

const pulse = keyframes`
    0% { transform: scale(1); }
    50% { transform: scale(1.1); }
    100% { transform: scale(1); }
`;

const AnimationDiv = styled.div`
    height: 50%;
    width: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    animation: ${pulse} 1s infinite;
`

function Loading(){
    return(
        <Container>
            <AnimationDiv>
                <Iconcomponent height={'25%'} fill={'var(--color-text-primary)'}/>
            </AnimationDiv>
        </Container>
    )
}

export default Loading;