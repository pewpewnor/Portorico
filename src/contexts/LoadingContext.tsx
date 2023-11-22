import { Dispatch, SetStateAction, createContext } from "react";

const LoadingContext = createContext<
	[boolean, Dispatch<SetStateAction<boolean>>]
>([false, () => {}]);

export default LoadingContext;
