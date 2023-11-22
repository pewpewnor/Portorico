"use client";

import Loading from "@/components/layouts/Loading";
import ErrorMessage from "@/components/ui/ErrorMessage";
import client from "@/lib/axios";
import cutString from "@/lib/cut";
import { templateNames } from "@/templates/templates";
import { Website } from "@/types/model";
import { Validations } from "@/types/responses";
import {
	Button,
	Card,
	CardBody,
	CardFooter,
	CardHeader,
	Divider,
	Input,
	Link,
	Modal,
	ModalBody,
	ModalContent,
	ModalFooter,
	ModalHeader,
	Select,
	SelectItem,
	useDisclosure,
} from "@nextui-org/react";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { ChangeEvent, useEffect, useState } from "react";
import { CiSearch } from "react-icons/ci";
import { MdDriveFileRenameOutline } from "react-icons/md";

type CreateWebsiteResponse = {
	validations?: Validations;
	website?: Website;
};

type UpdateWebsiteResponse = {
	validations?: Validations;
};

const defaultInputData = {
	name: "",
	templateName: "",
	description: "",
};

export default function DashboardPage() {
	const router = useRouter();

	const [isLoding, setIsLoading] = useState(true);
	const {
		isOpen: isOpenCreate,
		onOpen: onOpenCreate,
		onOpenChange: onOpenChangeCreate,
	} = useDisclosure();
	const {
		isOpen: isOpenEdit,
		onOpen: onOpenEdit,
		onOpenChange: onOpenChangeEdit,
	} = useDisclosure();
	const [inputSearch, setInputSearch] = useState("");
	const [websites, setWebsites] = useState<Website[]>([]);
	const [inputData, setInputData] = useState(defaultInputData);
	const [validations, setValidations] = useState<Validations>({});
	const [oldWebsiteId, setOldWebsiteId] = useState("");

	useEffect(() => {
		(async () => {
			setIsLoading(true);
			try {
				const res = await client.get("/authed/websites");
				const data = res.data as Website[];

				if (res.status === 400) {
					console.log("error 400");
				} else if (res.status === 200) {
					setWebsites(data);
				}
			} catch (error) {
				setIsLoading(false);
				router.replace("/login");
				return;
			}
			setIsLoading(false);
		})();
	}, [router]);

	function handleChange(event: ChangeEvent<HTMLInputElement>) {
		setInputData((prev) => ({
			...prev,
			[event.target.name]: event.target.value,
		}));
	}

	async function handleCreate(onClose: () => void) {
		setIsLoading(true);
		try {
			const res = await client.post("/authed/website", inputData);
			const data = res.data as CreateWebsiteResponse;

			if (res.status === 400 && data.validations) {
				setValidations(data.validations);
			} else if (res.status === 200 && data.website) {
				const website = data.website;
				setWebsites((prev) => [...prev, website]);
				setValidations({});
				onClose();
			}
		} catch (error) {
			setIsLoading(false);
			router.replace("/login");
			return;
		}
		setIsLoading(false);
	}

	async function handleUpdate(onClose: () => void) {
		setIsLoading(true);
		try {
			const res = await client.put("/authed/website", {
				...inputData,
				websiteId: oldWebsiteId,
			});
			const data = res.data as UpdateWebsiteResponse;

			if (res.status === 400 && data.validations) {
				setValidations(data.validations);
			} else if (res.status === 200) {
				setWebsites((prev) =>
					prev.map((website) =>
						website.id === oldWebsiteId
							? {
									...website,
									name: inputData.name,
									description: inputData.description,
							  }
							: website
					)
				);

				setValidations({});
				onClose();
			}
		} catch (error) {
			setIsLoading(false);
			router.replace("/login");
			return;
		}
		setIsLoading(false);
	}

	return (
		<div className="mt-8 flex flex-col gap-6 px-6 sm:px-14 xl:px-32">
			{isLoding && <Loading />}
			<div className="flex items-center justify-center gap-4">
				<p className="hidden whitespace-nowrap text-lg font-medium sm:block">
					My Websites
				</p>
				<Input
					classNames={{
						input: "text-md",
						inputWrapper: "h-10 bg-default-600/20",
					}}
					placeholder="Type to search..."
					size="md"
					startContent={<CiSearch className="w-8" />}
					type="search"
					value={inputSearch}
					onChange={(event: ChangeEvent<HTMLInputElement>) => {
						setInputSearch(event.target.value);
					}}
				/>
				<Button
					onPress={() => {
						setInputData(defaultInputData);
						onOpenCreate();
					}}
					className="w-36 bg-sky-blue"
				>
					Add New
				</Button>
			</div>
			{websites.length ? (
				<div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
					{websites.map(
						(website, index) =>
							(website.name.includes(inputSearch) ||
								website.description.includes(inputSearch)) && (
								<Card key={index} className="w-full">
									<CardHeader className="flex justify-between">
										<div className="flex gap-3">
											<Image
												alt="nextui logo"
												height={40}
												src="/favicon.ico"
												width={50}
											/>
											<div className="flex flex-col">
												<p className="text-md">
													{cutString(
														website.name,
														20
													)}
												</p>
												<p className="text-small text-default-500">
													{website.visitorsThisMonth}{" "}
													visitors this month
												</p>
											</div>
										</div>
										<div className="flex items-center justify-center gap-3">
											<button
												onClick={() => {
													setOldWebsiteId(website.id);
													setInputData(website);
													onOpenEdit();
												}}
											>
												<MdDriveFileRenameOutline className="h-6 w-6" />
											</button>
											<Button className="bg-indigo-600 font-medium text-white">
												Edit
											</Button>
										</div>
									</CardHeader>
									<Divider />
									<CardBody>
										<p>{website.description}</p>
									</CardBody>
									<Divider />
									<CardFooter className="flex justify-between">
										<Link
											href={"/p/" + website.name}
											className="text-indigo-600"
											showAnchorIcon
											underline="always"
										>
											View this website
										</Link>
										<p className="text-small text-default-500">
											{website.templateName}
										</p>
									</CardFooter>
								</Card>
							)
					)}
				</div>
			) : (
				<div className="flex items-center justify-center gap-2">
					<p className="text-center">
						You haven&apos;t created any website yet.
					</p>
					<Link
						href="/create"
						underline="always"
						className="text-indigo-600"
					>
						Create A New Website
					</Link>
				</div>
			)}

			{/* create new website modal */}
			<Modal
				isOpen={isOpenCreate}
				onOpenChange={onOpenChangeCreate}
				placement="top-center"
			>
				<ModalContent>
					{(onClose) => (
						<>
							<ModalHeader className="flex flex-col gap-1">
								Add New Website
							</ModalHeader>
							<ModalBody>
								<Input
									autoFocus
									label="website name"
									variant="bordered"
									name="name"
									value={inputData.name}
									onChange={handleChange}
								/>
								{inputData.name.length > 0 && (
									<p>
										Your website link would be:{" "}
										<p className="underline">
											portorico.io/p/
											{cutString(inputData.name, 40)}
										</p>
									</p>
								)}
								{validations.name && (
									<ErrorMessage message={validations.name} />
								)}
								<Input
									label="description"
									variant="bordered"
									name="description"
									value={inputData.description}
									onChange={handleChange}
								/>
								{validations.description && (
									<ErrorMessage
										message={validations.description}
									/>
								)}
								<Select
									isRequired
									label="select a template"
									variant="bordered"
									className="w-full"
									onChange={(
										event: ChangeEvent<HTMLSelectElement>
									) => {
										setInputData((prev) => ({
											...prev,
											templateName:
												templateNames[
													+event.target.value
												],
										}));
									}}
								>
									{templateNames.map(
										(templateName, index) => (
											<SelectItem key={index}>
												{templateName}
											</SelectItem>
										)
									)}
								</Select>
								{validations.templateName && (
									<ErrorMessage
										message={validations.templateName}
									/>
								)}
							</ModalBody>
							<ModalFooter>
								<Button
									color="danger"
									variant="flat"
									onPress={() => {
										setInputData(defaultInputData);
										setValidations({});
										onClose();
									}}
								>
									Cancel
								</Button>
								<Button
									color="primary"
									onPress={() => handleCreate(onClose)}
								>
									Create
								</Button>
							</ModalFooter>
						</>
					)}
				</ModalContent>
			</Modal>

			{/* change website information modal */}
			<Modal
				isOpen={isOpenEdit}
				onOpenChange={onOpenChangeEdit}
				placement="top-center"
			>
				<ModalContent>
					{(onClose) => (
						<>
							<ModalHeader className="flex flex-col gap-1">
								Edit Website Information
							</ModalHeader>
							<ModalBody>
								<Input
									autoFocus
									label="website name"
									variant="bordered"
									name="name"
									value={inputData.name}
									onChange={handleChange}
								/>
								{inputData.name.length > 0 && (
									<p>
										Your website link would be:{" "}
										<p className="underline">
											portorico.io/p/
											{cutString(inputData.name, 40)}
										</p>
									</p>
								)}
								{validations.name && (
									<ErrorMessage message={validations.name} />
								)}
								<Input
									label="description"
									variant="bordered"
									name="description"
									value={inputData.description}
									onChange={handleChange}
								/>
								{validations.description && (
									<ErrorMessage
										message={validations.description}
									/>
								)}
							</ModalBody>
							<ModalFooter>
								<Button
									color="danger"
									variant="flat"
									onPress={() => {
										setInputData(defaultInputData);
										setValidations({});
										onClose();
									}}
								>
									Cancel
								</Button>
								<Button
									color="primary"
									onPress={() => handleUpdate(onClose)}
								>
									Confirm
								</Button>
							</ModalFooter>
						</>
					)}
				</ModalContent>
			</Modal>
		</div>
	);
}
