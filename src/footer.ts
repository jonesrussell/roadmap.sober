// src/footer.ts
export function getFooterTemplate() {
    return `
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
      </footer>
    `;
  }
  