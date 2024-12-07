import styled from "styled-components";
import { useTranslation } from "react-i18next";
import { FaFaceSmileWink } from "react-icons/fa6";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const Container = styled.div`
    display: flex;
    width: 100%;
    height: 150vh;
    justify-content: center;
    align-items: center;
    flex-direction: column;
`;

const Title = styled.h1`
    width: 100%;
    text-align: center;
    font-weight: bolder;
    font-size: 2em;
    margin-bottom: 25px;
`;

const Button = styled.button`
    background-color: var(--color-primary);
    border: none;
    width: 20%;
    padding: 20px 10px;
    font-weight: bold;
    font-size: 1em;
    margin-inline: auto;
    border-radius: 10px;
    margin-block: 10px;
    cursor: pointer;
`;

const Icon = styled(FaFaceSmileWink)`
    font-size: 5em;
    margin-inline: auto;
    margin-bottom: 15px;
`;

export default function Home() {
    const { t, i18n } = useTranslation();
    const navigate = useNavigate();

    useEffect(() => {
        const language = localStorage.getItem('language');
        const theme = localStorage.getItem('theme');
        if (language) i18n.changeLanguage(language);
        if (theme) {
            document.body.classList.toggle('theme-white', theme === 'white');
        }
    }, [i18n]);

    const handleLanguageChange = (lang) => {
        i18n.changeLanguage(lang);
        localStorage.setItem('language', lang);
    };

    const toggleTheme = () => {
        const isWhiteMode = document.body.classList.toggle('theme-white');
        localStorage.setItem('theme', isWhiteMode ? 'white' : 'dark');
    };

    const goToLogin = () => {
        navigate('/login');
    };

    return (
        <Container>
            <Icon />
            <Title>{t("welcome")}</Title>
            <Button onClick={() => handleLanguageChange("pt")}>PT</Button>
            <Button onClick={() => handleLanguageChange("en")}>EN</Button>
            <Button onClick={toggleTheme}>{t("toggle_theme")}</Button>
            <Button onClick={goToLogin}>{t("go_to_login")}</Button>
        </Container>
    );
}