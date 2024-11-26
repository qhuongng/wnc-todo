"use client";

import { ChangeEvent } from "react";

import { filterTodo } from "@/lib/redux/todosSlice";
import { useAppDispatch } from "@/lib/hooks";

const TodoFilter: React.FC = () => {
    const dispatch = useAppDispatch();

    return (
        <input
            type="text"
            placeholder="Search for a to-do"
            className="input input-bordered w-full max-w-md"
            onChange={(e: ChangeEvent) => {
                if (e.target instanceof HTMLInputElement) {
                    dispatch(filterTodo(e.target.value));
                }
            }}
        />
    );
};

export default TodoFilter;
