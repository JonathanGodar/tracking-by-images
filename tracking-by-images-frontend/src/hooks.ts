export async function handle({ event, resolve }: any) {
  // const response = await resolve(event, {ssr: false});
	const response = await resolve(event)
  return response;
}
