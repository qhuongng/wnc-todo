"use client";

import { useSelector } from "react-redux";
import { RootState } from "@/lib/store";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

interface ProtectedRouteProps {
    children: React.ReactNode;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {

    const storedUser = useSelector((state: RootState) => state.user.username);
    const router = useRouter();

    useEffect(() => {
        if (storedUser === null) {
            router.push("/login");
        }
    }, [router]);

    return <div>{children}</div>;
};

export default ProtectedRoute;
