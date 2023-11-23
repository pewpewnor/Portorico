import Image from "next/image";
import "./style.css";

interface TestTemplateProps {
	isEditing: boolean;
	text1?: string;
	text2?: string;
}

export default function BasicLanding(props: TestTemplateProps) {
	const text1 = props.text1 ?? "MY PORTFOLIO";
	const text2 = props.text2 ?? "abour me lorem ipsum";

	return (
		<div>
			<div className="hero">
				<nav>
					<h2 className="logo">
						Porto<span>rico</span>
					</h2>
					<ul>
						<li>
							<a href="#">Home</a>
						</li>
						<li>
							<a href="#">About Us</a>
						</li>
						<li>
							<a href="#">Services</a>
						</li>
						<li>
							<a href="#">Skills</a>
						</li>
						<li>
							<a href="#">Contact Us</a>
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

			<section className="about">
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

			<div className="service">
				<div className="title">
					<h2>Our Services</h2>
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

			<div className="contact-me">
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
		</div>
	);
}
