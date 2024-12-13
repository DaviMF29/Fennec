import { useState, useEffect, useRef } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { useTranslation } from "react-i18next";
import './i18n';
import Home from "./pages/Home";
import Login from "./pages/Login";
import Search from "./pages/Search";
import Loading from "./components/Loading";

import LayoutContainer from './components/layout/LayoutContainer';
import LayoutTopBar from './components/layout/LayoutTopBar';
import LayoutMainArea from './components/layout/LayoutMainArea';
import LayoutSideNavbar from './components/layout/LayoutSideNavbar';
import LayoutSideContent from './components/layout/LayoutSideContent';
import LayoutRouteArea from './components/layout/LayoutRouteArea';
import LayoutBottomBar from './components/layout/LayoutBottomBar';

export default function App() {
  const { i18n } = useTranslation();
  const [isLoading, setIsLoading] = useState(true);

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

    const timer = setTimeout(() => setIsLoading(false), 2000);
    return () => clearTimeout(timer);
  }, [i18n]);

  if (isLoading) return <Loading />;

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
  const [WindowWidth, setWindowWidth] = useState(window.innerWidth);
  const [scrollDirection, setScrollDirection] = useState("down");
  const [lastScrollY, setLastScrollY] = useState(0);
  const layoutRouteAreaRef = useRef(null);

  useEffect(() => {
    const handleResize = () => {
      setWindowWidth(window.innerWidth);
    };

    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  useEffect(() => {
    const handleScroll = () => {
      if (!layoutRouteAreaRef.current) return;
      const currentScrollY = layoutRouteAreaRef.current.scrollTop;

      if (currentScrollY > lastScrollY) {
        setScrollDirection("up");
      } else {
        setScrollDirection("down");
      }
      setLastScrollY(currentScrollY);
    };

    const scrollableElement = layoutRouteAreaRef.current;
    scrollableElement?.addEventListener("scroll", handleScroll);

    return () => scrollableElement?.removeEventListener("scroll", handleScroll);
  }, [lastScrollY]);

  return (
    <LayoutContainer>
      <LayoutTopBar scrollDirection={scrollDirection}/>
      <LayoutMainArea>
        {WindowWidth > 600 && <LayoutSideNavbar scrollDirection={scrollDirection}/>}
        {WindowWidth > 600 && <LayoutSideContent />}
        <LayoutRouteArea ref={layoutRouteAreaRef}>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/search" element={<Search />} />
          </Routes>
        </LayoutRouteArea>
      </LayoutMainArea>
      {WindowWidth < 600 && <LayoutSideNavbar scrollDirection={scrollDirection}/>}
      <LayoutBottomBar style={{ display: WindowWidth <= 600 ? 'none' : 'flex' }}/>
    </LayoutContainer>
  );
}