import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Login from "./Pages/Login";
import UserPage from "./Pages/UserPage";
import LibAdminPage from "./Pages/LibAdminPage";
import AdminPage from "./Pages/AdminPage";

const App: React.FC = () => {
  return (
    <Router>
      <div className="d-flex">
        <div className="flex-grow-1">
          <Routes>
            <Route path="/" element={<Login></Login>} />
            <Route path="/user" element={<UserPage></UserPage>} />
            <Route path="/libadmin" element={<LibAdminPage></LibAdminPage>} />
            <Route path="/admin" element={<AdminPage></AdminPage>} />
          </Routes>
        </div>
      </div>
    </Router>
  );
};

export default App;
