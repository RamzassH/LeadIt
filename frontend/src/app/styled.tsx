import * as React from 'react';
import { GlobalStyles } from '@mui/material';

export default function GlobalStyle() {
    return (
        <React.Fragment>
            <GlobalStyles
                styles={{
        '@font-face': [
            {
                fontFamily: 'Roboto',
                src: "url('/fonts/Roboto.ttf') format('truetype'), url('/fonts/Roboto-Italic.ttf') format('truetype')",
                fontStyle: 'normal',
            },
        ],
        '* , ::before, ::after': {
            boxSizing: 'border-box',
                borderWidth: 0,
                borderStyle: 'solid',
                borderColor: '#e5e7eb',
        },
        '::before, ::after': {
            '--tw-content': '""',
        },
        html: {
            lineHeight: 1.5,
                WebkitTextSizeAdjust: '100%',
                MozTabSize: 4,
                tabSize: 4,
                fontFamily:
            'ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"',
                fontFeatureSettings: 'normal',
                fontVariationSettings: 'normal',
                WebkitTapHighlightColor: 'transparent',
        },
        body: {
            margin: 0,
                lineHeight: 'inherit',
        },
        hr: {
            height: 0,
                color: 'inherit',
                borderTopWidth: '1px',
        },
        'abbr[title]': {
            textDecoration: 'underline dotted',
        },
        'h1, h2, h3, h4, h5, h6': {
            fontSize: 'inherit',
                fontWeight: 'inherit',
        },
        a: {
            color: 'inherit',
                textDecoration: 'inherit',
        },
        'b, strong': {
            fontWeight: 'bolder',
        },
        'code, kbd, samp, pre': {
            fontFamily:
                'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace',
                    fontFeatureSettings: 'normal',
                fontVariationSettings: 'normal',
                fontSize: '1em',
        },
        small: {
            fontSize: '80%',
        },
        'sub, sup': {
            fontSize: '75%',
                lineHeight: 0,
                position: 'relative',
                verticalAlign: 'baseline',
        },
        sub: {
            bottom: '-0.25em',
        },
        sup: {
            top: '-0.5em',
        },
        table: {
            textIndent: 0,
                borderColor: 'inherit',
                borderCollapse: 'collapse',
        },
        'button, input, optgroup, select, textarea': {
            fontFamily: 'inherit',
                fontFeatureSettings: 'inherit',
                fontVariationSettings: 'inherit',
                fontSize: '100%',
                fontWeight: 'inherit',
                lineHeight: 'inherit',
                letterSpacing: 'inherit',
                color: 'inherit',
                margin: 0,
                padding: 0,
        },
        'button, select': {
            textTransform: 'none',
        },
        'button, input[type="button"], input[type="reset"], input[type="submit"]': {
            WebkitAppearance: 'button',
                backgroundColor: 'transparent',
                backgroundImage: 'none',
        },
        ':-moz-focusring': {
            outline: 'auto',
        },
        ':-moz-ui-invalid': {
            boxShadow: 'none',
        },
        progress: {
            verticalAlign: 'baseline',
        },
        '::-webkit-inner-spin-button, ::-webkit-outer-spin-button': {
            height: 'auto',
        },
        '[type="search"]': {
            WebkitAppearance: 'textfield',
                outlineOffset: '-2px',
        },
        '::-webkit-search-decoration': {
            WebkitAppearance: 'none',
        },
        '::-webkit-file-upload-button': {
            WebkitAppearance: 'button',
                font: 'inherit',
        },
        summary: {
            display: 'list-item',
        },
        'blockquote, dl, dd, h1, h2, h3, h4, h5, h6, hr, figure, p, pre': {
            margin: 0,
        },
        fieldset: {
            margin: 0,
                padding: 0,
        },
        legend: {
            padding: 0,
        },
        'ol, ul, menu': {
            listStyle: 'none',
                margin: 0,
                padding: 0,
        },
        dialog: {
            padding: 0,
        },
        textarea: {
            resize: 'vertical',
        },
        'input::placeholder, textarea::placeholder': {
            opacity: 1,
                color: '#9ca3af',
        },
        'button, [role="button"]': {
            cursor: 'pointer',
        },
        ':disabled': {
            cursor: 'default',
        },
        'img, svg, video, canvas, audio, iframe, embed, object': {
            display: 'block',
                verticalAlign: 'middle',
        },
        'img, video': {
            maxWidth: '100%',
                height: 'auto',
        },
        '[hidden]:where(:not([hidden="until-found"]))': {
            display: 'none',
        },
    }}
    />
    </React.Fragment>
);
}
/*
// Стилизованные компоненты
export const StyledButton = styled('button')({
    cursor: 'pointer',
    backgroundColor: 'transparent',
    backgroundImage: 'none',
    WebkitAppearance: 'button',
    fontFamily: 'inherit',
    fontFeatureSettings: 'inherit',
    fontVariationSettings: 'inherit',
    fontSize: '100%',
    fontWeight: 'inherit',
    lineHeight: 'inherit',
    letterSpacing: 'inherit',
    color: 'inherit',
    margin: 0,
    padding: 0,
    textTransform: 'none',
});

export const StyledInput = styled('input')({
    fontFamily: 'inherit',
    fontFeatureSettings: 'inherit',
    fontVariationSettings: 'inherit',
    fontSize: '100%',
    fontWeight: 'inherit',
    lineHeight: 'inherit',
    letterSpacing: 'inherit',
    color: 'inherit',
    margin: 0,
    padding: 0,
});

export const StyledTextarea = styled('textarea')({
    fontFamily: 'inherit',
    fontFeatureSettings: 'inherit',
    fontVariationSettings: 'inherit',
    fontSize: '100%',
    fontWeight: 'inherit',
    lineHeight: 'inherit',
    letterSpacing: 'inherit',
    color: 'inherit',
    margin: 0,
    padding: 0,
    resize: 'vertical',
});

export const StyledSelect = styled('select')({
    fontFamily: 'inherit',
    fontFeatureSettings: 'inherit',
    fontVariationSettings: 'inherit',
    fontSize: '100%',
    fontWeight: 'inherit',
    lineHeight: 'inherit',
    letterSpacing: 'inherit',
    color: 'inherit',
    margin: 0,
    padding: 0,
    textTransform: 'none',
});

export const StyledImg = styled('img')({
    display: 'block',
    verticalAlign: 'middle',
    maxWidth: '100%',
    height: 'auto',
});

export const StyledVideo = styled('video')({
    display: 'block',
    verticalAlign: 'middle',
    maxWidth: '100%',
    height: 'auto',
});

export const StyledTable = styled('table')({
    textIndent: 0,
    borderColor: 'inherit',
    borderCollapse: 'collapse',
});

export const StyledLink = styled('a')({
    color: 'inherit',
    textDecoration: 'inherit',
});

export const StyledList = styled('ul')({
    listStyle: 'none',
    margin: 0,
    padding: 0,
});

export const StyledListItem = styled('li')({
    listStyle: 'none',
    margin: 0,
    padding: 0,
});

export const StyledHeading = styled('h1')({
    fontSize: 'inherit',
    fontWeight: 'inherit',
});

export const StyledParagraph = styled('p')({
    margin: 0,
});

export const StyledBlockquote = styled('blockquote')({
    margin: 0,
});

export const StyledHr = styled('hr')({
    height: 0,
    color: 'inherit',
    borderTopWidth: '1px',
});

export const StyledCode = styled('code')({
    fontFamily:
        'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace',
    fontFeatureSettings: 'normal',
    fontVariationSettings: 'normal',
    fontSize: '1em',
});

export const StyledPre = styled('pre')({
    fontFamily:
        'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace',
    fontFeatureSettings: 'normal',
    fontVariationSettings: 'normal',
    fontSize: '1em',
    margin: 0,
});

export const StyledSmall = styled('small')({
    fontSize: '80%',
});

export const StyledSub = styled('sub')({
    fontSize: '75%',
    lineHeight: 0,
    position: 'relative',
    verticalAlign: 'baseline',
    bottom: '-0.25em',
});

export const StyledSup = styled('sup')({
    fontSize: '75%',
    lineHeight: 0,
    position: 'relative',
    verticalAlign: 'baseline',
    top: '-0.5em',
});

export const StyledFieldset = styled('fieldset')({
    margin: 0,
    padding: 0,
});

export const StyledLegend = styled('legend')({
    padding: 0,
});

export const StyledDialog = styled('dialog')({
    padding: 0,
});

export const StyledSummary = styled('summary')({
    display: 'list-item',
});

export const StyledProgress = styled('progress')({
    verticalAlign: 'baseline',
});

export const StyledAbbr = styled('abbr')({
    textDecoration: 'underline dotted',
});

export const StyledB = styled('b')({
    fontWeight: 'bolder',
});

export const StyledStrong = styled('strong')({
    fontWeight: 'bolder',
});

export const StyledKbd = styled('kbd')({
    fontFamily:
        'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace',
    fontFeatureSettings: 'normal',
    fontVariationSettings: 'normal',
    fontSize: '1em',
});

export const StyledSamp = styled('samp')({
    fontFamily:
        'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace',
    fontFeatureSettings: 'normal',
    fontVariationSettings: 'normal',
    fontSize: '1em',
});

export const StyledFigure = styled('figure')({
    margin: 0,
});

export const StyledEmbed = styled('embed')({
    display: 'block',
    verticalAlign: 'middle',
});

export const StyledObject = styled('object')({
    display: 'block',
    verticalAlign: 'middle',
});

export const StyledIframe = styled('iframe')({
    display: 'block',
    verticalAlign: 'middle',
});

export const StyledCanvas = styled('canvas')({
    display: 'block',
    verticalAlign: 'middle',
});

export const StyledAudio = styled('audio')({
    display: 'block',
    verticalAlign: 'middle',
});

export const StyledHidden = styled('div')({
    display: 'none',
});
*/