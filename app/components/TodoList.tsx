import { Todo } from "../lib/todos";
import { Icon } from "@iconify/react/dist/iconify.js";

interface TodoListProps {
    todos: Todo[];
    markTodo: (id: number) => void;
}

const TodoList: React.FC<TodoListProps> = ({ todos, markTodo }) => {
    return (
        <div className="flex flex-col pb-8 px-8">
            {todos.map((todo: Todo) => (
                <div
                    key={todo.id}
                    className="flex rounded-xl bg-slate-50 shadow-xl mb-4 px-4 items-center justify-between"
                >
                    <p
                        className={`font-semibold text-lg text-wrap ${
                            todo.completed ? "text-slate-400 line-through" : "text-slate-800"
                        }`}
                    >
                        {todo.content}
                    </p>

                    {!todo.completed && (
                        <button
                            className="btn btn-square btn-primary ml-8"
                            onClick={() => markTodo(todo.id)}
                        >
                            <Icon icon="ic:round-check" style={{ fontSize: "24px" }} />
                        </button>
                    )}
                </div>
            ))}
        </div>
    );
};

export default TodoList;
