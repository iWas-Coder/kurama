import clsx from 'clsx';
import Link from '@docusaurus/Link';
import Heading from '@theme/Heading';
import styles from './index.module.css';
import logo from '@site/static/img/logo.png';

function TwoColumns({c1, c2, reverse}) {
  return (
    <div className={`TwoColumns ${reverse ? 'reverse' : ''}`}>
      <div className={`column first ${reverse ? 'right' : 'left'}`}>
        {c1}
      </div>
      <div className={`column last ${reverse ? 'left' : 'right'}`}>
        {c2}
      </div>
    </div>
  );
}

function Logo() {
  return (
    <img
      className={ styles.Logo }
      src={ logo } />
  );
}

function Title() {
  return (
    <Heading as='h1' className='hero__title'>
      Kurama
    </Heading>
  );
}

function Subtitle() {
  return (
    <p className='hero__subtitle'>
      Cloud-native job scheduler and orchestrator
    </p>
  );
}

function Button({isPrimary, txt, path}) {
  return (
    <div className={ styles.Button }>
    <Link className={`button button--${ isPrimary ? 'primary' : 'secondary' } button--lg`} to={ path }>
        { txt }
      </Link>
    </div>
  );
}

function ButtonPair() {
  return (
    <div className={ styles.ButtonPair }>
      <Button
        isPrimary
        txt='Get started'
        path='/docs/getting-started/2.1-getting-started'
      />
      <Button
        txt='Learn the basics >'
        path='/docs/overview/1.1-what-is-kurama'
      />
    </div>
  );
}

export default function HomepageHeader() {
  return (
    <header className={ clsx('hero hero--secondary', styles.heroBanner) }>
      <div className='container'>
        <TwoColumns
          reverse
          c1={ <Logo /> }
          c2={
            <>
              <Title />
              <Subtitle />
              <ButtonPair />
            </>
          }
        />
      </div>
    </header>
  );
}
