"use client";

import { useAppSelector, useAppDispatch } from "@/lib/hooks";
import { RootState } from "@/lib/store";
import { markTodo } from "@/lib/redux/todosSlice";

const TodoList = () => {
    const todos = useAppSelector((state: RootState) => state.todos.filteredTodos);
    const dispatch = useAppDispatch();

    return (
        <div className="flex flex-col pb-8 px-8">
            {todos.map((todo) => (
                <div
                    key={todo.id}
                    className="flex rounded-xl bg-slate-50 shadow-xl mb-4 px-4 items-center justify-between"
                >
                    <p
                        className={`font-semibold text-lg text-wrap ${todo.completed ? "text-slate-400 line-through" : "text-slate-800"
                            }`}
                    >
                        {todo.content}
                    </p>

                    <input
                        type="checkbox"
                        className="checkbox checkbox-lg checkbox-primary ml-4"
                        checked={todo.completed}
                        onChange={() => dispatch(markTodo(todo.id))}
                    />
                </div>
            ))}
        </div>
    );
};

export default TodoList;
