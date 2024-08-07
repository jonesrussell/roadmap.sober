import './style.css';

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <header>
    <h1>Start Here</h1>
    <nav>
      <a href="#community">Community</a> | 
      <a href="#login">Login</a> | 
      <a href="#signup">Sign Up</a>
    </nav>
  </header>
  <main>
    <h2>Sobriety Roadmaps</h2>
    <p>Are you just starting your journey to sobriety? Here are some beginner friendly roadmaps you should start with.</p>
    <h3>Understanding Addiction</h3>
    <p>Learn about the nature of addiction, its causes, and effects. Understand why it's not just a "bad habit".</p>
    <h3>Getting Help</h3>
    <p>Explore different resources available for help, such as therapy, support groups, and medication.</p>
    <h3>Maintaining Sobriety</h3>
    <p>Learn strategies to maintain sobriety in the long term, including dealing with triggers and cravings, and building a supportive social network.</p>
    <p>There is also a roadmap for family and friends of people dealing with addiction, which includes understanding their loved one's struggle and knowing how to provide effective support.</p>
    <h2>Tips for Beginners</h2>
    <p>Starting a journey to sobriety can be overwhelming, here are some tips to help you get started:</p>
    <ul>
      <li>Reach out for help: Don't try to do it alone.</li>
      <li>Take it one day at a time: Don't worry about the future, just focus on staying sober today.</li>
    </ul>
  </main>
<footer>
    <div class="footer-row">
      <p>© 2024 Sober Journey</p>
    </div>
    <div class="footer-row">
      <div class="footer-column">
        <h4>About</h4>
        <ul>
          <li><a href="#mission">Our Mission</a></li>
          <li><a href="#team">Our Team</a></li>
          <li><a href="#contact">Contact Us</a></li>
        </ul>
      </div>
      <div class="footer-column">
        <h4>Resources</h4>
        <ul>
          <li><a href="#roadmaps">Sobriety Roadmaps</a></li>
          <li><a href="#tips">Tips for Beginners</a></li>
          <li><a href="#community">Join the Community</a></li>
        </ul>
      </div>
    </div>
  </footer>`
