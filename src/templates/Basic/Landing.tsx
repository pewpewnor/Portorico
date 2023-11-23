import {
	Button,
	Input,
	ModalBody,
	ModalContent,
	ModalFooter,
	ModalHeader,
	Modal as NextModal,
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
}

export default function BasicLanding(props: TestTemplateProps) {
	const text1 = props.text1 ?? "My Name";
	const text2 = props.text2 ?? "about me lorem ipsum";

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
			<div className="hero">
				<nav>
					<h2 className="logo">
						<span>
							<button onClick={handleChange("text1", text1)}>
								{text1}
							</button>
						</span>
					</h2>
					<ul>
						<li>
							<a href="#home">Home</a>
						</li>
						<li>
							<a href="#aboutMe">About Me</a>
						</li>
						<li>
							<a href="#services">Services</a>
						</li>
						<li>
							<a href="#contactUs">Contact Us</a>
						</li>
					</ul>
					<a href="#" className="btn">
						Subscribe
					</a>
				</nav>

				<div className="content">
					<h4>Hello, my name is</h4>
					<h1>{text1}</h1>
					<h3>I&apos;m a Web Developer</h3>
					<div className="newslatter">
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

			<section className="about" id="aboutMe">
				<div className="main">
					<Image
						src="/template/basic/profile.jpg"
						width={500}
						height={500}
						alt="profile"
					/>
					<div className="about-text">
						<h2>About me</h2>
						<h5>Web Developer</h5>
						<p>{text2}</p>
						<button>Let&apos;s Talk</button>
					</div>
				</div>
			</section>

			<div className="service" id="services">
				<div className="title">
					<h2>My Services</h2>
				</div>

				<div className="box">
					<div className="card">
						<i className="fa-solid fa-bars"></i>
						<h5>Web Development</h5>
						<div className="pra">
							<p>
								Lorem ipsum dolor sit amet, consectetur
								adipiscing elit. Vestibulum tincidunt risus ac
								augue auctor accumsan.
							</p>
							<p className="text-center">
								<a className="button" href="#">
									Read more
								</a>
							</p>
						</div>
					</div>

					<div className="card">
						<i className="fa-regular fa-user"></i>
						<h5>Web Development</h5>
						<div className="pra">
							<p>
								Lorem ipsum dolor sit amet, consectetur
								adipiscing elit. Vestibulum tincidunt risus ac
								augue auctor accumsan.
							</p>
							<p className="text-center">
								<a className="button" href="#">
									Read more
								</a>
							</p>
						</div>
					</div>

					<div className="card">
						<i className="fa-regular fa-bell"></i>
						<h5>Web Development</h5>
						<div className="pra">
							<p>
								Lorem ipsum dolor sit amet, consectetur
								adipiscing elit. Vestibulum tincidunt risus ac
								augue auctor accumsan.
							</p>
							<p className="text-center">
								<a className="button" href="#">
									Read more
								</a>
							</p>
						</div>
					</div>
				</div>
			</div>

			<div className="contact-me" id="contactUs">
				<p>Let me make you a website</p>
				<a className="button-two" href="#">
					Hire Me
				</a>
			</div>

			<footer>
				<p>Robert Robertson</p>
				<p>
					Lorem ipsum dolor sit amet, consectetur adipiscing elit.
					Vestibulum tincidunt risus ac augue auctor accumsan.
				</p>
				<div className="social">
					<a href="#">
						<i className="fa-brands fa-instagram"></i>
					</a>
					<a href="#">
						<i className="fa-brands fa-facebook"></i>
					</a>
					<a href="#">
						<i className="fa-brands fa-linkedin"></i>
					</a>
				</div>
				<p className="end">Portorico</p>
			</footer>

			<NextModal
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
							<ModalFooter>
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
									}}
								>
									Create
								</Button>
							</ModalFooter>
						</>
					)}
				</ModalContent>
			</NextModal>
		</div>
	);
}
