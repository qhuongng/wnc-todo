import { Todo } from "../lib/todos";
import { Icon } from "@iconify/react/dist/iconify.js";

interface TodoItemProps {
    todo: Todo;
    markTodo: (id: number) => void;
}

const TodoItem: React.FC<TodoItemProps> = ({ todo, markTodo }) => {
    return (
        <div className="flex rounded-xl bg-slate-50 shadow-xl mb-4 px-4 items-center justify-between">
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
    );
};

export default TodoItem;
