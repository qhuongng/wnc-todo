"use client";

import { useSelector, useDispatch } from "react-redux";
import { RootState, AppDispatch } from "@/lib/store";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

interface ProtectedRouteProps {
    children: React.ReactNode;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {

    const storedUser = useSelector((state: RootState) => state.user.username);
    const router = useRouter();
    const dispatch = useDispatch<AppDispatch>();

    useEffect(() => {
        if (storedUser === null) {
            router.push("/login");
        }
    }, [router, dispatch]);

    return <div>{children}</div>;
};

export default ProtectedRoute;
