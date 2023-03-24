import { HashRouter, Navigate, Route, Routes } from "react-router-dom";
import Repositories from "../pages/repositories"
import Tags from "../pages/tags"

export default function AppRouter(): JSX.Element {
    return (
        <HashRouter>
            <Routes>
                <Route path="/" element={<Navigate to="repositories" />} />
                <Route path="/repositories" element={<Repositories />} />
                <Route path="/tags" element={<Tags />} />
            </Routes>
        </HashRouter>
    )
}
