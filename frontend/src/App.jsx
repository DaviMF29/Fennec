import { BrowserRouter, Routes, Route } from "react-router-dom";
import { useEffect } from "react";
import { useTranslation } from "react-i18next";
import './i18n';
import Home from "./pages/Home";
import Login from "./pages/Login";

import LayoutContainer from './components/layout/LayoutContainer';
import LayoutTopBar from './components/layout/LayoutTopBar';
import LayoutMainArea from './components/layout/LayoutMainArea';
import LayoutSideNavbar from './components/layout/LayoutSideNavbar';
import LayoutSideContent from './components/layout/LayoutSideContent';
import LayoutRouteArea from './components/layout/LayoutRouteArea';
import LayoutBottomBar from './components/layout/LayoutBottomBar';

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
        <Route path="/login" element={<Login />} />
        <Route path="*" element={<AppLayout />} />
      </Routes>
    </BrowserRouter>
  );
}

function AppLayout() {
  return (
    <LayoutContainer>
      <LayoutTopBar />
      <LayoutMainArea>
        <LayoutSideNavbar/>
        <LayoutSideContent />
        <LayoutRouteArea>
          <Routes>
            <Route path="/" element={<Home />} />
          </Routes>
        </LayoutRouteArea>
      </LayoutMainArea>
      <LayoutBottomBar />
    </LayoutContainer>
  );
}