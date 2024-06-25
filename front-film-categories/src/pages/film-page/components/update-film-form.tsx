import { format } from "date-fns"
import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import { useForm } from "react-hook-form"
import { useFilms } from "../../../hooks/films"
import { BaseFormProps } from "../../../types"
import { Button } from "../../../components/ui/button"
import { Input } from "../../../components/ui/input"
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "../../../components/ui/form"
import { Popover, PopoverContent, PopoverTrigger } from "../../../components/ui/popover"
import { cn } from "../../../lib/utils"
import { CalendarIcon } from "lucide-react"
import { Calendar } from "../../../components/ui/calendar"

interface UpdateFilmFormProps extends BaseFormProps {
    id: string;
    name: string;
    author: string;
    description: string;
    duration: number;
    release_date: Date;
}

const formSchema = z.object({
    name: z.string().min(1, {
        message: "Name is required",
    }).max(150, {
        message: "Name is too long",
    }),
    author: z.string().min(1, {
        message: "Author is required",
    }).max(150, {
        message: "Name is too long",
    }),
    description: z.string().max(150, {
        message: "Description is too long",
    }),
    duration: z.coerce.number().int().positive(),
    release_date: z.date(),
})

export function UpdateFilmForm({ onClose, onRefresh, id, name, author, description, duration, release_date }: UpdateFilmFormProps) {
    const { updateFilm } = useFilms();

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: name,
            author: author,
            description: description,
            duration: duration,
            release_date: new Date(release_date),
        },
    })

    async function onSubmit(values: z.infer<typeof formSchema>) {
        await updateFilm(id, values);
        onRefresh();
        if (onClose) { onClose() }
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                <FormField
                    control={form.control}
                    name="name"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Name</FormLabel>
                            <FormControl>
                                <Input placeholder="" {...field} />
                            </FormControl>
                            <FormDescription>
                                Name of the film.
                            </FormDescription>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="author"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Author</FormLabel>
                            <FormControl>
                                <Input placeholder="" {...field} />
                            </FormControl>
                            <FormDescription>
                                Name of the movie author.
                            </FormDescription>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="description"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Description</FormLabel>
                            <FormControl>
                                <Input placeholder="" {...field} />
                            </FormControl>
                            <FormDescription>
                                Description of the film
                            </FormDescription>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="duration"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Duration</FormLabel>
                            <FormControl>
                                <Input placeholder="" {...field} />
                            </FormControl>
                            <FormDescription>
                                Duration of the film in minutes.
                            </FormDescription>
                            <FormMessage />
                        </FormItem>
                    )}
                />

                <FormField
                    control={form.control}
                    name="release_date"
                    render={({ field }) => (
                        <FormItem className="flex flex-col">
                            <FormLabel>Release Date</FormLabel>
                            <Popover>
                                <PopoverTrigger asChild>
                                    <FormControl>
                                        <Button
                                            variant={"outline"}
                                            className={cn(
                                                "w-full pl-3 text-left font-normal",
                                                !field.value && "text-muted-foreground"
                                            )}
                                        >
                                            {field.value ? (
                                                format(field.value, "PPP")
                                            ) : (
                                                <span>Pick a date</span>
                                            )}
                                            <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                                        </Button>
                                    </FormControl>
                                </PopoverTrigger>
                                <PopoverContent className="w-auto p-0" align="start">
                                    <Calendar
                                        mode="single"
                                        selected={field.value}
                                        onSelect={field.onChange}
                                        disabled={(date) =>
                                            date > new Date() || date < new Date("1900-01-01")
                                        }
                                        initialFocus
                                    />
                                </PopoverContent>
                            </Popover>
                            <FormDescription>
                                Date when the film was released.
                            </FormDescription>
                            <FormMessage />
                        </FormItem>
                    )}
                />

                <div className="flex justify-between">
                    <Button type="button" variant="ghost" onClick={onClose}>Close</Button>
                    <Button type="submit">Update</Button>
                </div>
            </form>
        </Form>
    )
}