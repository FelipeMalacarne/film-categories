import { Navbar } from "./components/navbar";
import { Outlet } from "react-router-dom";
import { Button } from "./components/ui/button";
import { useFilms } from "./hooks/films";
import { useEffect } from "react";

function App() {
  const { films, createFilm, isLoading } = useFilms();

  useEffect(() => {
    console.log(films);
  }, [films]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <main>
        <Button
          onClick={() =>
            createFilm({
              name: "New Film",
              duration: 120,
              description: "New Film Description",
              author: "New Film Author",
              release_date: new Date(),
            })
          }
        >
          Create Film
        </Button>
        <Navbar />
        <Outlet />
      </main>
    </>
  );
}

export default App;
