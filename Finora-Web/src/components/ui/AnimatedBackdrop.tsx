export function AnimatedBackdrop() {
  return (
    <>
      <div className='market-glow -left-20 top-10 h-64 w-64 bg-neon-purple' />
      <div className='market-glow right-10 top-1/3 h-72 w-72 bg-neon-cyan [animation-delay:1.4s]' />
      <div className='market-glow bottom-10 left-1/3 h-80 w-80 bg-neon-green [animation-delay:2.2s]' />
      <div className='ai-grid-bg absolute inset-0 opacity-55' />
    </>
  )
}
