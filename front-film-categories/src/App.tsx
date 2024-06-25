import { Navbar } from "./components/navbar";
import { Outlet } from "react-router-dom";

function App() {
  return (
    <>
    <main className="h-screen flex flex-col">
        <Navbar />
        <Outlet />
      </main>
    </>
  );
}

export default App;
