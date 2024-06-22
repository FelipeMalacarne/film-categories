import axios from "axios";
import { BaseURL } from "./config";
import { useEffect } from "react";
import { Navbar } from "./components/navbar";
import { Outlet } from "react-router-dom";

function App() {
  const getFilms = async () => {
    const response = await axios.get(BaseURL + "/film");
    const data = await response.data;
    console.log(data);
  };

  useEffect(() => {
    getFilms();
  }, []);

  return (
    <>
      <main>
        <Navbar />
        <Outlet />
      </main>
    </>
  );
}

export default App;
