// src/main.ts
import { getHeaderTemplate } from './header.ts';
import { getFooterTemplate } from './footer.ts';
import './roadmap.sober.ts';
import './style.css';

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  ${getHeaderTemplate()}
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
  ${getFooterTemplate()}
`;
