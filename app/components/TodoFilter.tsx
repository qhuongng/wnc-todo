import React, { ChangeEvent } from "react";

interface TodoFilterProps {
    filterTodo: (keyword: string) => void;
}

const TodoFilter: React.FC<TodoFilterProps> = ({ filterTodo }) => {
    return (
        <input
            type="text"
            placeholder="Search for a to-do"
            className="input input-bordered w-full max-w-xl"
            onChange={(e: ChangeEvent) => {
                if (e.target instanceof HTMLInputElement) {
                    filterTodo(e.target.value);
                }
            }}
        />
    );
};

export default TodoFilter;
