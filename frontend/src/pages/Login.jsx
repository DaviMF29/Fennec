import { useState, useEffect } from 'react';
import styled from "styled-components";
import Logo from '../assets/iconProvisorio.png';
import { useTranslation } from "react-i18next";
import { Link } from 'react-router-dom';
import api from '../services/api';
import { useNavigate } from 'react-router-dom';

function Login() {
    const { t } = useTranslation();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [username, setUsername] = useState('');
    const [repeatPassword, setRepeatPassword] = useState('');
    const [isRegister, setIsRegister] = useState(true);
    const [isLightTheme, setIsLightTheme] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        const theme = localStorage.getItem('theme');
        setIsLightTheme(theme === 'white');
    }, []);


    const handleSubmit = async (event) => {
        event.preventDefault();
        document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'none';

        if (isRegister) {
            if (!email || !password || !username || !repeatPassword) {
                document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorField');
                document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                return;
            }
            if (!email.includes('@') || !email.includes('.')) {
                document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorEmail');
                document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                return;
            }
            if (username.length < 3) {
                document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorUsername');
                document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                return;
            }
            if (password.length < 6) {
                document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorPasswordSize');
                document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                return;
            }
            if (password !== repeatPassword) {
                document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorPasswordMismatch');
                document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                return;
            }

            try {
                const data = {
                    birth_date: 'null',
                    created_at: new Date().toISOString(),
                    email,
                    name: username,
                    password,
                    updated_at: new Date().toISOString(),
                    username,
                };

                const response = await api.post('/user/', data);

                if (response.status === 200) {
                    console.log(response.data.message)
                    localStorage.setItem('token', response.data.message || '');
                    navigate('/')
                }
            } catch (error) {
                if(!error.response){
                    document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorBadrequest');
                    document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                    console.log(error)
                    return;
                } else{
                    if (error.response.status === 400) {
                        document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorField');
                        document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                    }
                    if (error.response.status === 409) {
                        document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorInUse');
                        document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                    }
                    else{
                        document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorBadrequest');
                        document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                        return;
                    }
                }
            }
        } else {
            if (!email || !password) {
                document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorField');
                document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                return;
            }

            try {
                const data = { email, password };

                const response = await api.post('/login', data);

                if (response.status === 200) {
                    console.log(response.data)
                    localStorage.setItem('token', response.data.token || '');
                    navigate('/')
                }
            } catch (error){
                if (error.response.status === 401) {
                    document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_invalidCredentials');
                    document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                    return;
                } else{
                    document.querySelector(`.${ErrorLabel.styledComponentId}`).innerHTML = t('LOGIN_errorBadrequest');
                    document.querySelector(`.${ErrorLabel.styledComponentId}`).style.display = 'block';
                    return;
                }
            }
        }
    };

    const handleToggleRegister = () => setIsRegister(!isRegister);

    return (
        <BaseDiv>
            <LinkComponent to={'/'}>
                <LogoImage src={Logo} alt="Logo" $isLightTheme={isLightTheme} />
            </LinkComponent>
            <Container>
                <LinkComponentMobile to={'/'}>
                    <LogoImageMobile src={Logo} alt="Logo" $isLightTheme={isLightTheme}/>
                </LinkComponentMobile>
                <Title>{t("welcome")}</Title>
                <Subtitle>
                    {isRegister ? t("already_have_account") : t("new_here")}{' '}
                    <Span onClick={handleToggleRegister}>
                        {isRegister ? t("login") : t("create_account")}
                    </Span>
                </Subtitle>
                <Form onSubmit={handleSubmit}>
                    {isRegister && (
                        <Input
                            type="text"
                            placeholder={t("username_placeholder")}
                            value={username}
                            onChange={(e) => setUsername(e.target.value.toLowerCase().replace(/\s/g, ''))}
                        />
                    )}
                    <Input
                        type="text"
                        placeholder={t("email_placeholder")}
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <Input
                        type="password"
                        placeholder={t("password_placeholder")}
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    {isRegister && (
                        <Input
                            type="password"
                            placeholder={t("repeat_password_placeholder")}
                            value={repeatPassword}
                            onChange={(e) => setRepeatPassword(e.target.value)}
                        />
                    )}
                    <ErrorLabel>{t("fill_all_fields")}</ErrorLabel>
                    <Button type="submit">{isRegister ? t("register") : t("login")}</Button>
                    <HelpText>{t("login_issues")}</HelpText>
                </Form>
            </Container>
        </BaseDiv>
    );
}

const Container = styled.div`
    display: flex;
    flex-direction: column;
    align-items: start;
    justify-content: center;
    padding: 40px 50px;
    border-radius: 10px;
    text-align: left;
    min-width: 350px;
    max-width: 350px;
    margin-right: auto;
    margin-block: auto;
    background-color: var(--color-primary);

    @media (max-width: 1000px) {
        margin-inline: auto;
    }

    @media (max-width: 490px) {
        min-width: 0px;
        width: 90%;
        padding: 30px 20px;
    }
`;

const Title = styled.h1`
    font-size: 1.5em;
    font-weight: bolder;
    margin: 0px 0 10px;
`;

const Subtitle = styled.p`
    font-size: 0.8rem;
    font-weight: bold;
    margin-bottom: 30px;

    span {
        color: var(--color-wombat);
        cursor: pointer;
        transition: all 0.3s;

        &:hover {
            color: var(----color-text-primary);
        }
    }
`;

const Form = styled.form`
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
`;

const Input = styled.input`
    width: calc(100% - 40px);
    font-size: 0.9em;
    font-weight: bold;
    color: var(--color-text-primary);
    background-color: var(--color-secondary);
    border-radius: 10px;
    border: 1px solid var(--color-secondary);
    padding: 15px 20px;
    margin: 0 20px 15px;
    outline: none;
    transition: all 0.3s;

    &:hover, &:focus {
        border-color: var(--color-text-primary);
    }
`;

const ErrorLabel = styled.h4`
    color: rgb(221, 36, 36);
    font-size: 0.75em;
    width: 100%;
    text-align: left;
    display: none;
`;

const Button = styled.button`
    background-color: var(--color-wombat);
    border-radius: 10px;
    width: 100%;
    padding: 15px;
    font-size: 1em;
    color: var(--color-text-primary);
    font-weight: bolder;
    border: none;
    outline: none;
    margin: 20px 0;
    transition: all 0.3s;

    &:hover, &:focus {
        background-color: var(--color-text-primary);
        color: var(--color-background);
        cursor: pointer;
    }
`;

const HelpText = styled.h3`
    font-size: 0.65em;
    width: 100%;
    text-align: left;
    transition: all 0.3s;
    opacity: 0.7;

    &:hover {
        color: var(--color-wombat);
        cursor: pointer;
    }
`;

const Span = styled.span`
    color: var(--color-wombat);
`

const BaseDiv = styled.div`
    display: flex;
    flex-direction: row;
    width: 100%;
`;

const LogoImage = styled.img`
    width: 23vw;
    height: 23vw;
    filter: ${(props) => (props.isLightTheme ? 'invert(95%)' : 'none')};
`;

const LogoImageMobile = styled.img`
    height: 50px;
    margin-right: auto;
    margin-inline: auto;
    display: none;
    margin-bottom: 25px;

    @media (max-width: 1000px) {
        display: flex;
        filter: ${(props) => (props.isLightTheme ? 'invert(95%)' : 'none')};
    }
`;

const LinkComponent = styled(Link)`
    margin: 0;
    padding: 0;
    height: auto;
    margin-block: auto;
    width: auto;
    margin-inline: auto;

    @media (max-width: 1000px) {
        display: none;
    }
`

const LinkComponentMobile = styled(Link)`
    margin: 0;
    padding: 0;
    height: auto;
    margin-block: auto;
    width: auto;
    margin-inline: auto;
`

export default Login;