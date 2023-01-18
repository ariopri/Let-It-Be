import { Outlet, Route, Routes } from 'react-router-dom';
import RoutesOutlet from './utils/RoutesOutlet';
import LandingPage from './pages/01-LandingPage/LandingPage';
import Informasi from './pages/02-InformasiPage/Informasi';
import About from './pages/04-Tentang/About';
import HubungiKami from './pages/03-Hubungi/Hubungi';
import Modul from './pages/05-Modul/Modul';
import Register from './pages/07-Daftar/Register';
import CaraMendafatarPengajar from './pages/08-DaftarPengajar/CaraDaftarPengajar';
import Protected from './protected/protectedRoute';
import Login from './pages/06-LoginPage/Login';
import ModulDetail from './pages/09-ModulDetail/ModulDetail';
import SocialProfileSimple from './pages/10-ProfilePage/Profile';
import NotFound from './pages/11-NotFound/NotFound';
import Event from './pages/12-Event/Event';

function App() {
  return (
    <Routes>
      <Route path="/" element={<RoutesOutlet />}>
        <Route index element={<LandingPage />} />
        <Route path="informasi" element={<Informasi />} />
        <Route path="hubungi" element={<HubungiKami />} />
        <Route path="tentang" element={<About />} />
        <Route element={<Protected />}>
          <Route path="modul" element={<Outlet />}>
            <Route index element={<Modul />} />
            <Route path=":name" element={<Outlet />}>
              <Route path=":id" element={<ModulDetail />} />
            </Route>
          </Route>
        </Route>
        <Route path="event" element={<Event />} />
        <Route path="profile" element={<SocialProfileSimple />} />
        <Route path="masuk" element={<Login />} />
        <Route path="mendaftar" element={<Register />} />
        <Route
          path="cara_mendaftar_instruktur"
          element={<CaraMendafatarPengajar />}
        />
      </Route>
      <Route path="*" element={<NotFound />} />
    </Routes>
  );
}

export default App;
