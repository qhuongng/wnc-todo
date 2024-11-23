import { Icon } from "@iconify/react/dist/iconify.js";

import { markTodo } from "@/lib/redux/todosSlice";
import { useAppDispatch } from "@/lib/hooks";
import { Todo } from "@/lib/todos";

interface TodoItemProps {
    todo: Todo;
}

const TodoItem: React.FC<TodoItemProps> = ({ todo }) => {
    const dispatch = useAppDispatch();

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
                    onClick={() => dispatch(markTodo(todo.id))}
                >
                    <Icon icon="ic:round-check" style={{ fontSize: "24px" }} />
                </button>
            )}
        </div>
    );
};

export default TodoItem;
