"use client";

import { useAppSelector, useAppDispatch } from "@/lib/hooks";
import { RootState } from "@/lib/store";
import { fetchTodos, updateTodo } from "@/lib/redux/todosApi";
import { useEffect } from "react";

const TodoList = () => {
    const todos = useAppSelector((state: RootState) => state.todos.filteredTodos);
    const status = useAppSelector((state: RootState) => state.todos.status);
    const dispatch = useAppDispatch();

    useEffect(() => {
        dispatch(fetchTodos());
    }, [dispatch]);

    if (status === "failure")
        return (
            <p className="flex items-center justify-center">Đã xảy ra lỗi kết nối với máy chủ!</p>
        );

    return (
        <div className="flex flex-col pb-8 px-8">
            {status === "pending" ? (
                <div className="flex flex-col space-y-4 px-5">
                    <div className="skeleton w-full h-32" />
                    <div className="skeleton w-full h-32" />
                </div>
            ) : todos.length > 0 ? (
                todos.map((todo) => (
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
                            onChange={() => dispatch(updateTodo(todo))}
                        />
                    </div>
                ))
            ) : (
                <p className="flex items-center justify-center">Chưa có việc cần làm!</p>
            )}
        </div>
    );
};

export default TodoList;
