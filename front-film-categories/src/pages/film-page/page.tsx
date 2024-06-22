import { useFilms } from "../../hooks/films";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "../../components/ui/table";
import { Button } from "../../components/ui/button";

export default function FilmPage() {
  const { films, createFilm } = useFilms();

  return (
    <div className="space-y-8 flex flex-col">
      <div className="p-8">
        <h1 className="p-4 text-2xl font-semibold border-b">Filmes</h1>
        <Button
          variant={"outline"}
          onClick={() =>
            createFilm({
              name: "Filme 1",
              duration: 120,
              release_date: new Date(),
            })
          }
        >
          Criar
        </Button>
      </div>
      <div className="flex-1">
        <Table>
          <TableCaption>A list of your recent invoices.</TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Author</TableHead>
              <TableHead>Release Date</TableHead>
              <TableHead className="text-right">Duration</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {films ? (
              films.map((film) => (
                <TableRow key={film.id}>
                  <TableCell>{film.name}</TableCell>
                  <TableCell>{film.author}</TableCell>
                  <TableCell>{film.release_date.toLocaleString()}</TableCell>
                  <TableCell className="text-right">{film.duration}</TableCell>
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell colSpan={4}>Loading...</TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
