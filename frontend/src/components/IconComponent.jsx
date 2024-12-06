/* eslint-disable react/prop-types */

function SvgComponent({ width, height, fill }) {
    return (
        <svg
            width={width}
            height={height}
            viewBox="0 0 675 675"
            fill={fill}
            xmlns="http://www.w3.org/2000/svg"
        >
            <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M161.038 197.356L19.262 20.389v300.523S262.382 504.4 337 654.611V270s-176-4.552-176 128c0-90.752.038-200.644.038-200.644zM513.962 197.356L655.738 20.389v300.523S412.618 504.4 338 654.611V270s176-4.552 176 128c0-90.752-.038-200.644-.038-200.644z"
                fill={fill}
                stroke={fill}
                strokeWidth={20}
                strokeLinejoin="round"
            />
        </svg>
    )
}

export default SvgComponent