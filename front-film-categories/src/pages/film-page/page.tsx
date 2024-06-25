import { useState } from "react";
import { useFilms } from "../../hooks/films";
import { Button } from "../../components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "../../components/ui/table";
import { AddFilmForm } from "./components/add-film-form";
import { UpdateFilmForm } from "./components/update-film-form";
import PopUpDialog from "../../components/pop-up-dialog";
import { Skeleton } from "../../components/ui/skeleton";
import { SquarePen, Trash2 } from "lucide-react";

type OpenDialogs = {
    [key: string]: boolean;
};

export default function FilmPage() {
    const { films, isLoading, deleteFilm, getFilms } = useFilms();
    const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
    const [openDialogs, setOpenDialogs] = useState<OpenDialogs>({});

    const setIsUpdateDialogOpen = (id: string, isOpen: boolean) => {
        setOpenDialogs((prev) => ({ ...prev, [id]: isOpen }));
    };


    const handleCreateDialogClose = () => {
        setIsCreateDialogOpen(false);
    };

    return (
        <div className="flex-1 space-y-4 p-8 pt-6">
            <div className="flex items-center justify-between space-y-2">
                <h2 className="text-3xl font-bold tracking-tight">Films</h2>

                <PopUpDialog
                    isOpen={isCreateDialogOpen}
                    onOpenChange={setIsCreateDialogOpen}
                    title="Add Film"
                    text="Add a new Film"
                    FormComponent={
                        <AddFilmForm
                            onClose={handleCreateDialogClose}
                            onRefresh={getFilms}
                        />
                    }
                />
            </div>
            <Table className="w-full">
                <TableHeader>
                    <TableRow>
                        <TableHead className="hidden lg:table-cell">ID</TableHead>
                        <TableHead>Name</TableHead>
                        <TableHead>Author</TableHead>
                        <TableHead className="hidden lg:table-cell max-h-[100px]">Description</TableHead>
                        <TableHead>Duration</TableHead>
                        <TableHead>Release Date</TableHead>
                        <TableHead>Category</TableHead>
                        <TableHead>Created At</TableHead>
                        <TableHead className="text-left">Actions</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {films && (
                        films.map((film) => (
                            <TableRow key={film.id}>
                                <TableCell className="hidden lg:table-cell">{film.id}</TableCell>
                                <TableCell>{film.name}</TableCell>
                                <TableCell>{film.author}</TableCell>
                                <TableCell className="hidden lg:table-cell max-h-[100px]">
                                    {film.description}
                                </TableCell>
                                <TableCell>
                                    {film.duration + " min"}
                                </TableCell>
                                <TableCell>{new Date(film.release_date).toLocaleString('en-US', {
                                    year: 'numeric',
                                    month: 'long',
                                    day: 'numeric'
                                }
                                )}</TableCell>
                                <TableCell>{film.category?.name}</TableCell>
                                <TableCell>{new Date(film.created_at).toLocaleString('pt-BR')}</TableCell>
                                <TableCell className="text-left flex items-center">
                                    <PopUpDialog
                                        isOpen={openDialogs[film.id] || false}
                                        onOpenChange={(isOpen) => setIsUpdateDialogOpen(film.id, isOpen)}
                                        title="Update Film"
                                        text=<SquarePen/>
                                        FormComponent={
                                            <UpdateFilmForm
                                                onClose={() => setIsUpdateDialogOpen(film.id, false)}
                                                onRefresh={getFilms}
                                                id={film.id}
                                                name={film.name}
                                                author={film.author}
                                                description={film.description}
                                                duration={film.duration}
                                                release_date={film.release_date}
                                            />
                                        }
                                    />
                                    <Button
                                        variant="destructive"
                                        onClick={() => deleteFilm(film.id)}
                                        className="ml-2"
                                    >
                                        <Trash2/>
                                    </Button>
                                </TableCell>
                            </TableRow>
                        ))
                    )}
                    {!films && isLoading && (
                        <TableRow>
                            {Array.from({ length: 8 }).map((_, index) => (
                                <TableCell colSpan={1} key={index}>
                                    <Skeleton className="h-4" />
                                </TableCell>
                            ))}
                        </TableRow>
                    )}

                    {!films && !isLoading && (
                        <TableRow>
                            <TableCell colSpan={8} className="text-center">
                                No films found
                            </TableCell>
                        </TableRow>
                    )}

                </TableBody>
            </Table>
        </div>
    );
}
