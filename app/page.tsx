"use client";

import { useState } from "react";

import TodoForm from "@/components/TodoForm";
import TodoList from "@/components/TodoList";
import TodoFilter from "@/components/TodoFilter";

import Todos, { Todo } from "@/lib/todos";

const Home = () => {
    const [todos, setTodos] = useState<Todo[]>(Todos());
    const [filteredTodos, setFilteredTodos] = useState<Todo[]>(Todos());

    const addTodo = (content: string) => {
        const newTodo: Todo = {
            id: todos.length + 1,
            content: content,
            completed: false,
        };

        const updatedTodos = [...todos, newTodo];

        setTodos(updatedTodos);
        setFilteredTodos(updatedTodos);
    };

    const markTodo = (id: number) => {
        const updatedTodos = todos.map((todo) => {
            if (todo.id === id) {
                return { ...todo, completed: !todo.completed };
            }

            return todo;
        });

        setTodos(updatedTodos);
        setFilteredTodos(updatedTodos);
    };

    const filterTodo = (keyword: string) => {
        setTodos(
            filteredTodos.filter((todo) => {
                return todo.content.toLowerCase().includes(keyword.toLowerCase());
            })
        );
    };

    return (
        <article className="prose lg:prose-md md:prose-md sm:prose-xs">
            <div className="flex p-12 h-screen w-screen justify-center">
                <div className="flex flex-col bg-accent rounded-xl h-full shadow-xl w-1/2 overflow-auto">
                    <div className="flex w-full rounded-xl p-8 sticky-header justify-between align-middle">
                        <TodoFilter filterTodo={filterTodo} />
                        <TodoForm addTodo={addTodo} />
                    </div>

                    <TodoList todos={todos} markTodo={markTodo} />
                </div>
            </div>
        </article>
    );
};

export default Home;
