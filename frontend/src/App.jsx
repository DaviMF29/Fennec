import { BrowserRouter, Routes, Route } from "react-router-dom";
import { useEffect } from "react";
import { useTranslation } from "react-i18next";
import './i18n';
import Home from "./pages/Home";
import Welcome from "./pages/Login";

export default function App() {
  const { i18n } = useTranslation();

  useEffect(() => {
    const savedLanguage = localStorage.getItem('language');
    const savedTheme = localStorage.getItem('theme');

    if (savedLanguage) {
      i18n.changeLanguage(savedLanguage);
    } else {
      i18n.changeLanguage('pt');
    }

    if (savedTheme) {
      document.body.classList.toggle('theme-white', savedTheme === 'white');
    }
  }, [i18n]);

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/welcome" element={<Welcome />} />
      </Routes>
    </BrowserRouter>
  );
}