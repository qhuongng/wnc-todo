// Components that interact with the Redux store need to be client components
// Accessing the store requires React context, which is only available in client components

"use client";

import { useRef } from "react";
import { Provider } from "react-redux";

import { makeStore, AppStore } from "@/lib/store";
import { fetchTodosSuccess } from "@/lib/redux/todosSlice";

import Todos from "@/lib/todos";

const StoreProvider = ({ children }: { children: React.ReactNode }) => {
    const storeRef = useRef<AppStore>();

    if (!storeRef.current) {
        // Create the store instance the first time this renders
        storeRef.current = makeStore();
        storeRef.current.dispatch(fetchTodosSuccess(Todos()));
    }

    return <Provider store={storeRef.current}>{children}</Provider>;
};

export default StoreProvider;
