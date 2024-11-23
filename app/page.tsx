"use client";

import TodoForm from "@/components/TodoForm";
import TodoList from "@/components/TodoList";
import TodoFilter from "@/components/TodoFilter";

const Home = () => {
    return (
        <article className="prose lg:prose-md md:prose-md sm:prose-xs">
            <div className="flex p-12 h-screen w-screen justify-center">
                <div className="flex flex-col bg-accent rounded-xl h-full shadow-xl w-1/2 overflow-auto">
                    <div className="flex w-full rounded-xl p-8 sticky-header justify-between align-middle">
                        <TodoFilter />
                        <TodoForm />
                    </div>

                    <TodoList />
                </div>
            </div>
        </article>
    );
};

export default Home;
