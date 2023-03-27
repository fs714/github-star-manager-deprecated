import { HashRouter, Navigate, Route, Routes } from 'react-router-dom';

import AppLayout from '../layout';
import Page403 from '../pages/errorPages/403';
import Page404 from '../pages/errorPages/404';
import Repositories from '../pages/repositories';
import Tags from '../pages/tags';

export default function AppRouter(): JSX.Element {
  return (
    <HashRouter>
      <Routes>
        <Route path="/" element={<AppLayout />}>
          <Route path="/" element={<Repositories />} />
          <Route path="/repositories" element={<Repositories />} />
          <Route path="/tags" element={<Tags />} />
          <Route path="403" element={<Page403 />} />
          <Route path="404" element={<Page404 />} />
          <Route path="*" element={<Navigate to="404" />} />
        </Route>
      </Routes>
    </HashRouter>
  );
}
