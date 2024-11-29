export interface Todo {
    id: number;
    content: string;
    completed: boolean;
}

const Todos = () => [
    {
        id: 1,
        content: "Negotiate a peace treaty with the intergalactic squirrel overlords",
        completed: false,
    },
    {
        id: 2,
        content: "Invent a time machine to steal the Mona Lisa from the Louvre in 1850",
        completed: false,
    },
];

export default Todos;
