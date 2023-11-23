import {
	Button,
	Input,
	Modal,
	ModalBody,
	ModalContent,
	ModalFooter,
	ModalHeader,
	useDisclosure,
} from "@nextui-org/react";
import Image from "next/image";
import { ChangeEvent, useState } from "react";
import "./style.css";

interface TestTemplateProps {
	isEditing: boolean;
	changeField: (fieldName: string, value: string) => void;
	text1?: string;
	text2?: string;
	text3?: string;
	text4?: string;
	text5?: string;
	text6?: string;
	text7?: string;
	text8?: string;
	text9?: string;
	text10?: string;
	text11?: string;
	text12?: string;
	text13?: string;
	text14?: string;
	text15?: string;
}

export default function BasicLanding(props: TestTemplateProps) {
	const text1 = props.text1 ?? "My Name";
	const text2 = props.text2 ?? "Hello there,";
	const text3 = props.text3 ?? "My Name is";
	const text4 =
		props.text4 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text5 =
		props.text5 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text6 =
		props.text6 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text7 =
		props.text7 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text8 =
		props.text8 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text9 =
		props.text9 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text10 =
		props.text10 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text11 =
		props.text11 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text12 =
		props.text12 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text13 =
		props.text13 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text14 =
		props.text14 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";
	const text15 =
		props.text15 ??
		"Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus, dolorum!";

	const { isOpen, onOpen, onOpenChange } = useDisclosure();
	const [fieldName, setFieldName] = useState("");
	const [fieldValue, setFieldValue] = useState("");

	function handleChange(fieldName: string, fieldValue: string) {
		return () => {
			setFieldName(fieldName);
			setFieldValue(fieldValue);
			onOpen();
		};
	}

	return (
		<div id="home">
			<div className="basic_hero">
				<div className="basic_nav">
					<p className="basic_logo basic_h2">
						<span
							onClick={handleChange("text1", text1)}
							className="cursor-pointer"
						>
							{text1}
						</span>
					</p>
					<ul className="basic_ul">
						<li className="basic_li">
							<a href="#home" className="basic_a">
								Home
							</a>
						</li>
						<li className="basic_li">
							<a href="#aboutMe" className="basic_a">
								About Me
							</a>
						</li>
						<li className="basic_li">
							<a href="#services" className="basic_a">
								Services
							</a>
						</li>
						<li className="basic_li">
							<a href="#contactUs" className="basic_a">
								Contact Us
							</a>
						</li>
					</ul>
					<a href="#" className="basic_btn">
						Subscribe
					</a>
				</div>

				<div className="basic_content">
					<p
						className="basic_h4 cursor-pointer"
						onClick={handleChange("text2", text2)}
					>
						{text2}
					</p>
					<p
						className="basic_h1 cursor-pointer"
						onClick={handleChange("text3", text3)}
					>
						{text3}
					</p>
					<p
						className="basic_h3 cursor-pointer"
						onClick={handleChange("text4", text4)}
					>
						{text4}
					</p>
					<div className="basic_newslatter">
						<form>
							<input
								type="email"
								name="email"
								id="mail"
								placeholder="Enter Your Email"
							/>
							<input
								type="submit"
								name="submit"
								value="Lets Start"
							/>
						</form>
					</div>
				</div>
			</div>

			<section className="basic_about" id="aboutMe">
				<div className="basic_main">
					<Image
						src="/template/basic/profile.jpg"
						width={500}
						height={500}
						alt="profile"
					/>
					<div className="basic_about-text">
						<h2>About me</h2>
						<h5
							className="cursor-pointer"
							onClick={handleChange("text5", text5)}
						>
							{text5}
						</h5>
						<p
							className="cursor-pointer"
							onClick={handleChange("text6", text6)}
						>
							{text6}
						</p>
						<button>Let&apos;s Talk</button>
					</div>
				</div>
			</section>

			<div className="basic_service" id="services">
				<div className="basic_title">
					<h2>My Services</h2>
				</div>

				<div className="basic_box">
					<div className="basic_card">
						<i className="basic_fa-solid basic_fa-bars"></i>
						<p
							className="cursor-pointer"
							onClick={handleChange("text15", text15)}
						>
							{text15}
						</p>
						<div className="basic_pra">
							<p
								className="cursor-pointer"
								onClick={handleChange("text7", text7)}
							>
								{text7}
							</p>
							<p className="text-center">
								<a className="basic_button" href="#">
									Read more
								</a>
							</p>
						</div>
					</div>

					<div className="basic_card">
						<i className="basic_fa-solid basic_fa-bars"></i>
						<p
							className="cursor-pointer"
							onClick={handleChange("text8", text8)}
						>
							{text8}
						</p>
						<div className="basic_pra">
							<p
								className="cursor-pointer"
								onClick={handleChange("text9", text9)}
							>
								{text9}
							</p>
							<p className="basic_text-center">
								<a className="basic_button" href="#">
									Read more
								</a>
							</p>
						</div>
					</div>

					<div className="basic_card">
						<i className="basic_fa-regular basic_fa-bell"></i>
						<p
							className="cursor-pointer"
							onClick={handleChange("text10", text10)}
						>
							{text10}
						</p>
						<div className="basic_pra">
							<p
								className="cursor-pointer"
								onClick={handleChange("text11", text11)}
							>
								{text11}
							</p>
							<p className="text-center">
								<a className="basic_button" href="#">
									Read more
								</a>
							</p>
						</div>
					</div>
				</div>
			</div>

			<div className="basic_contact-me" id="contactUs">
				<p
					className="cursor-pointer"
					onClick={handleChange("text12", text12)}
				>
					{text11}
				</p>
				<a className="button-two" href="#">
					Hire Me
				</a>
			</div>

			<div className="basic_footer">
				<p
					className="cursor-pointer"
					onClick={handleChange("text13", text13)}
				>
					{text13}
				</p>
				<p
					className="cursor-pointer"
					onClick={handleChange("text14", text14)}
				>
					{text14}
				</p>
				<div className="basic_social">
					<a href="#" className="basic_a">
						<i className="basic_fa-brands basic_fa-instagram"></i>
					</a>
					<a href="#" className="basic_a">
						<i className="basic_fa-brands basic_fa-facebook"></i>
					</a>
					<a href="#" className="basic_a">
						<i className="basic_fa-brands basic_fa-linkedin"></i>
					</a>
				</div>
				<p className="basic_end">{text1}</p>
			</div>

			<Modal
				isOpen={isOpen}
				onOpenChange={onOpenChange}
				placement="top-center"
			>
				<ModalContent>
					{(onClose) => (
						<>
							<ModalHeader className="flex flex-col gap-1">
								Change {fieldName}
							</ModalHeader>
							<ModalBody>
								<Input
									autoFocus
									label={fieldName}
									variant="bordered"
									name="name"
									value={fieldValue}
									onChange={(
										event: ChangeEvent<HTMLInputElement>
									) => {
										setFieldValue(event.target.value);
									}}
								/>
							</ModalBody>
							<ModalFooter className="min-w-min">
								<Button
									color="danger"
									variant="flat"
									onPress={onClose}
								>
									Cancel
								</Button>
								<Button
									color="primary"
									onPress={() => {
										props.changeField(
											fieldName,
											fieldValue
										);
										onClose();
									}}
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
