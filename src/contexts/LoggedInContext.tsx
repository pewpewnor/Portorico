import { Dispatch, SetStateAction, createContext } from "react";

const LoggedInContext = createContext<[boolean, () => void]>([false, () => {}]);

export default LoggedInContext;
