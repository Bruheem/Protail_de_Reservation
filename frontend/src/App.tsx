//
import { BrowserRouter as Router } from "react-router-dom";
import AppRoutes from "./routes";
import { UserProvider } from "./utils/AuthContext";

const App = () => {
  return (
    <>
      <Router>
        <UserProvider>
          <AppRoutes />
        </UserProvider>
      </Router>
    </>
  );
};

export default App;
