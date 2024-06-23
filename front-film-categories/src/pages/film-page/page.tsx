import { useFilms } from "../../hooks/films";
import { Button } from "../../components/ui/button";
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "../../components/ui/dialog";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "../../components/ui/table";
import { AddFilmForm } from "./components/add-film-form";

export default function FilmPage() {
  const { films, createFilm, updateFilm, deleteFilm } = useFilms();

  return (
    <div className="space-y-8 flex flex-col">
      <div className="p-8 flex items-center">
        <h1 className="p-4 text-2xl font-semibold border-b">Films</h1>
        <Dialog>
          <DialogTrigger asChild className=" ml-4">
            <Button variant="default">Add</Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[425px]">
            <DialogHeader>
              <DialogTitle>Add film</DialogTitle>
            </DialogHeader>
            <DialogContent>
              <AddFilmForm />
            </DialogContent>
            <DialogFooter>
              <Button type="submit">Save changes</Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>
      <div className="flex items-center justify-center">
        <div className="w-full max-w-screen-md mx-auto">
          <Table className="w-full">
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Author</TableHead>
                <TableHead>Release Date</TableHead>
                <TableHead>Duration</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {films ? (
                films.map((film) => (
                  <TableRow key={film.id}>
                    <TableCell>{film.name}</TableCell>
                    <TableCell>{film.author}</TableCell>
                    <TableCell>{film.release_date.toLocaleString()}</TableCell>
                    <TableCell>{film.duration}</TableCell>
                    {/* seria uma boa ideia botar svg aq */}
                    <TableCell>
                      <Button
                        className="mr-2"
                        variant="outline"
                        onClick={() => { }}
                      >
                        Editar
                      </Button>
                      <Button
                        variant="destructive"
                        onClick={() => deleteFilm(film.id)}
                      >
                        Excluir
                      </Button>
                    </TableCell>
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
    </div>
  );

}
