# Source Library by Era — Starting List

Draft starting list for the four witness agents' curated libraries, plus the Konbit, a forward-looking fifth agent planned for Phase 3. Each witness retrieves only from its own era's shelf, plus the shared cross-era shelf. Have a Haitian historian or a librarian review it before ingestion.

**Tags to store on every document:** era · language (FR / KR / EN) · type (primary / scholarship / testimony / journalism / notes) · rights (see below) · perspective notes.

---

## Phase 1: Start here (Witnesses 1 and 2)

Build and test the full pipeline with the two witnesses that have the deepest public-domain material. Witnesses 3 and 4 come later.

**Witness 1 (to 1804), public domain**
- Code Noir (1685)
- Moreau de Saint-Méry, *Description… de la partie française de l'isle Saint-Domingue* (1797), on Gallica
- Sonthonax's emancipation proclamation (1793) and the Convention's abolition decree (1794)
- Toussaint Louverture's letters and proclamations; the Constitution of 1801
- Haitian Declaration of Independence (1804)
- Colonial records: ANOM, John Carter Brown Library

**Witness 2 (1804–1915), public domain**
- Imperial Constitution of 1805
- Ordinance of Charles X (1825)
- Boyer's Code Rural (1826)
- Thomas Madiou, *Histoire d'Haïti*, on Gallica
- Beaubrun Ardouin, *Études sur l'histoire d'Haïti*, on Gallica (tag its elite and mulâtre perspective)
- Anténor Firmin, *De l'égalité des races humaines* (1885)
- Haitian laws and newspapers from the period, on dLOC

**Plus, for both:** open-access articles and dissertations, and your own notes on the copyrighted books (see the rights strategy below).

---

## Rights strategy

**Rights tags (store per document and per chunk):**
- `pd-us` / `pd-fr`: public domain. US: published through 1930 (as of 2026), plus all US federal government documents. France: author died more than 70 years ago. A work can be public domain in one country and still copyrighted in the other.
- `public-doc`: publicly released institutional documents (UN, Haitian government, audits). Generally fine to use; check each source's terms.
- `open-access`: openly licensed articles, theses, repositories. Record the license (CC BY, etc.).
- `archive-terms`: material under an archive's own terms of use (for example, Radio Haïti at Duke).
- `copyrighted-private`: full text of copyrighted books in private, personal use only. This is a legal gray area; not legal advice.
- `notes`: your own summaries of copyrighted works, written in your own words, with book and page citations.

**Approach**
1. Ingest public-domain and public documents first.
2. Add open-access scholarship (dLOC, university repositories, open journal articles).
3. Cover copyrighted books with notes: short summaries of their arguments, with citations. The witnesses benefit from the scholarship without ingesting full texts.
4. Full-text ingestion of copyrighted books is optional, private only, and always tagged so it can be removed with one query if the project ever goes public.

---

**Where to get the files**
- **Gallica (Bibliothèque nationale de France):** public-domain French texts (Moreau de Saint-Méry, Madiou, Ardouin, Firmin)
- **Digital Library of the Caribbean (dLOC):** Haitian newspapers, laws, government documents
- **Internet Archive / HathiTrust:** 19th- and early-20th-century books, US government documents
- **Archives nationales d'outre-mer (ANOM, Aix-en-Provence):** colonial administrative records for Saint-Domingue
- **John Carter Brown Library (Brown University):** colonial and revolutionary-era printed material
- **Duke University, Rubenstein Library:** the Radio Haïti archive (audio in Kreyòl and French, 1957–2003)
- **In Haiti:** Bibliothèque Haïtienne des Frères de l'Instruction Chrétienne; Bibliothèque Haïtienne des Spiritains

---

## Shared shelf (all witnesses)
- Michel-Rolph Trouillot, *Silencing the Past* (1995): how the history of the Revolution was recorded, and silenced. Scholarship, EN, ©
- Michel-Rolph Trouillot, *Haiti: State Against Nation* (1990): the state and the peasantry from 1804 to the Duvaliers. Scholarship, EN, ©
- Laurent Dubois, *Haiti: The Aftershocks of History* (2012): overview from 1804 to 2010. Scholarship, EN, ©

---

## Witness 1: Colony and Revolution (to 1804)

**Primary sources**
- Code Noir (1685). FR, PD
- Moreau de Saint-Méry, *Description… de la partie française de l'isle Saint-Domingue* (1797). FR, PD
- Sonthonax's emancipation proclamation (29 Aug 1793) and the National Convention's abolition decree (4 Feb 1794). FR, PD
- Toussaint Louverture's letters and proclamations; the Constitution of 1801. FR, PD
- Haitian Declaration of Independence (1 Jan 1804). FR, PD
- Bois Caïman (1791) accounts. The accounts were written later and are contested, so tag this as "disputed evidence."

**Scholarship**
- C.L.R. James, *The Black Jacobins* (1938). EN, ©
- Carolyn Fick, *The Making of Haiti: The Saint Domingue Revolution from Below* (1990). EN, ©
- Jean Fouchard, *Les Marrons de la liberté* (1972). A Haitian historian on marronage. FR, ©
- Laurent Dubois, *Avengers of the New World* (2004). EN, ©
- David Geggus, *Haitian Revolutionary Studies* (2002). EN, ©
- John Garrigus, *Before Haiti* (2006): free people of color in Saint-Domingue. EN, ©
- Philippe Girard, *Toussaint Louverture: A Revolutionary Life* (2016). EN, ©
- Sudhir Hazareesingh, *Black Spartacus* (2020). EN, ©

---

## Witness 2: Early Nation and the Debt (1804–1915)

**Primary sources**
- Imperial Constitution of 1805. FR, PD
- Royal ordinance of Charles X, 17 Apr 1825 (the indemnity). FR, PD
- Boyer's Code Rural (1826). FR, PD
- Thomas Madiou, *Histoire d'Haïti* (1847–48). FR, PD
- Beaubrun Ardouin, *Études sur l'histoire d'Haïti* (1853–60). FR, PD. Has an elite and mulâtre perspective; tag it.
- Anténor Firmin, *De l'égalité des races humaines* (1885). FR, PD

**Scholarship**
- Julia Gaffield, *Haitian Connections in the Atlantic World* (2015): recognition and diplomacy. EN, ©
- Marlene Daut, *Tropics of Haiti* (2015). EN, ©
- Johnhenry Gonzalez, *Maroon Nation* (2019): land and labor after independence. EN, ©
- Brenda Gayle Plummer, *Haiti and the Great Powers, 1902–1915* (1988). EN, ©
- Peter James Hudson, *Bankers and Empire* (2017): National City Bank and the December 1914 gold removal. Overlaps with Witness 3. EN, ©

**Journalism**
- *The New York Times*, "The Ransom" series (2022): what the indemnity cost Haiti. EN, ©

---

## Witness 3: Occupation to the Duvaliers (1915–1986) — Phase 2

**Primary sources**
- Haitian-American Convention (treaty of 1915). EN/FR, PD
- US Senate, *Inquiry into Occupation and Administration of Haiti and Santo Domingo* (hearings, 1921–22). EN, PD (US government)
- James Weldon Johnson, "Self-Determining Haiti," *The Nation* (1920). EN, PD in the US
- *Foreign Relations of the United States* (FRUS) volumes on Haiti. EN, PD (US government)
- Radio Haïti archive (Duke), Duvalier-era broadcasts. KR/FR, archive terms
- Amnesty International reports on the Duvalier regime. EN, ©

**Scholarship**
- Hans Schmidt, *The United States Occupation of Haiti, 1915–1934* (1971). EN, ©
- Mary Renda, *Taking Haiti* (2001). EN, ©
- Matthew J. Smith, *Red and Black in Haiti* (2009): politics from 1934 to 1957. EN, ©
- Bernard Diederich & Al Burt, *Papa Doc* (1969). EN, ©
- Elizabeth Abbott, *Haiti: The Duvaliers and Their Legacy* (1988). EN, ©

**Haitian voices of the era**
- Jean Price-Mars, *Ainsi parla l'oncle* (1928): the indigéniste response to the occupation. FR. Likely PD in the US (published before 1931) but still under copyright in France (author died in 1969).

---

## Witness 4: 1986 to Today — Phase 2

**Primary sources**
- Constitution of 1987 (and the 2011 amendments). FR/KR, public-doc
- UN Security Council resolutions: MINUSTAH (Res. 1542, 2004), BINUH (Res. 2476, 2019), Multinational Security Support mission (Res. 2699, 2023), and later resolutions. EN/FR, public-doc
- BINUH and UN human rights (OHCHR) reports on Haiti. EN/FR, public-doc
- UN statements on the cholera epidemic, including the Secretary-General's 2016 apology. EN, public-doc
- Government of Haiti, Post-Disaster Needs Assessment (2010 earthquake). FR/EN, public-doc
- Cour Supérieure des Comptes, PetroCaribe audit reports (2019–20). FR, public-doc
- Radio Haïti archive (to 2003). KR/FR, archive terms

**Scholarship and long-form journalism**
- Amy Wilentz, *The Rainy Season* (1989): the end of the Duvalier era and its aftermath. EN, ©
- Paul Farmer, *The Uses of Haiti* (1994; later editions). EN, ©
- Robert Fatton Jr., *Haiti's Predatory Republic* (2002) and *The Roots of Haitian Despotism* (2007): stress internal governance. EN, ©
- Alex Dupuy, *The Prophet and Power* (2007): Aristide, critical. EN, ©
- Peter Hallward, *Damming the Flood* (2007): Aristide, sympathetic. Pair it with Dupuy for balance. EN, ©
- Jonathan Katz, *The Big Truck That Went By* (2013): the earthquake and aid. EN, ©
- Jake Johnston, *Aid State* (2024). EN, ©

**Live sources (allowlist for web search; this witness only)**
- Le Nouvelliste (FR), AyiboPost (KR/FR), The Haitian Times (EN)
- Reuters, AP, the *Miami Herald*'s Haiti coverage
- International Crisis Group, BINUH and OHCHR updates

---

## The Konbit: "How can we move from here" — Phase 3

A fifth, forward-looking agent, named after the Haitian tradition of collective work. It helps people in Haiti and in the diaspora think through how to build or contribute something: a school, a business, a nonprofit, an investment.

**Design decisions**
- **Separate agent, not the orchestrator.** The orchestrator stays a neutral router (who speaks, handoffs, turn limits) and gives no advice.
- **Separate from Witness 4.** Witnesses testify about what happened; the Konbit advises. Keeping the roles apart stops advice from leaking into testimony.
- **Doesn't hold the witnesses' libraries.** It pulls history through the same `invite_witness` handoff tool. Example: "Witness 2, why did land ownership become so fragile?" or "Witness 4, what happened to housing projects after the earthquake?"
- **Not era-bound.** It may hand off to any witness.

**Its own library (practical, current)**
- Lessons from aid and development: what worked and what failed, especially after 2010 (Katz, *The Big Truck That Went By*; Johnston, *Aid State*; Farmer, *The Uses of Haiti*; evaluations of post-earthquake projects)
- Diaspora contribution: remittance flows and data (for example, the World Bank), hometown associations, diaspora investment initiatives
- Legal and administrative basics: registering a business (Ministry of Commerce and Industry) and a nonprofit (Ministry of Planning and External Cooperation), land title and property issues. Verify current procedures before relying on them.
- A vetted directory of reputable local organizations, by sector and region

**Live sources (allowlist)**
- Witness 4's news allowlist, plus current security information (BINUH and OHCHR reports, government travel advisories) and current regulations

**Guardrails**
- Clearly not a legal, financial or security adviser; points users to qualified professionals and local partners.
- Honest about current risks, especially security, without discouraging people from contributing.
- Favors partnering with local organizations over imported solutions, the clearest lesson of the aid era.
- Separates what the sources show from its own suggestions, and says when the evidence is thin.

---

## Balance notes
- Explanations of "how Haiti got here" split roughly into **external causes** (the indemnity, occupation, foreign intervention, aid) and **internal causes** (predatory governance, elite capture, dictatorship). Each era's shelf should include both views, tagged as such, so witnesses can present disagreement honestly.
- Haitian authors are underrepresented in English-language scholarship. Add more Haitian historians writing in French and Kreyòl as the library grows; ask reviewers for recommendations.
- Tag disputed evidence (for example, Bois Caïman or partisan memoirs) so a witness says "accounts differ" instead of stating it as fact.

---

## Later: outreach to historians
Contact contemporary Haitian historians through their public academic emails to:
- review the source list for each era and fill the gaps in Haitian, French-language and Kreyòl scholarship
- point to open-access versions of their own work, or grant permission to ingest it
- check the witness personas and the "disputed evidence" tags

Not needed for Phase 1. Worth doing before Witnesses 3 and 4, and before any public release.
