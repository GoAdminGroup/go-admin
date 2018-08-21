/*
	Ethereal by HTML5 UP
	html5up.net | @ajlkn
	Free for personal and commercial use under the CCA 3.0 license (html5up.net/license)
*/

(function($) {

	// Settings.
		var settings = {

			// Keyboard shortcuts.
				keyboardShortcuts: {

					// If true, enables scrolling via keyboard shortcuts.
						enabled: true,

					// Sets the distance to scroll when using the left/right arrow keys.
						distance: 50

				},

			// Scroll wheel.
				scrollWheel: {

					// If true, enables scrolling via the scroll wheel.
						enabled: true,

					// Sets the scroll wheel factor. (Ideally) a value between 0 and 1 (lower = slower scroll, higher = faster scroll).
						factor: 1

				},

			// Scroll zones.
				scrollZones: {

					// If true, enables scrolling via scroll zones on the left/right edges of the scren.
						enabled: true,

					// Sets the speed at which the page scrolls when a scroll zone is active (higher = faster scroll, lower = slower scroll).
						speed: 15

				},

			// Dragging.
				dragging: {

					// If true, enables scrolling by dragging the main wrapper with the mouse.
						enabled: true,

					// Sets the momentum factor. Must be a value between 0 and 1 (lower = less momentum, higher = more momentum, 0 = disable momentum scrolling).
						momentum: 0.875,

					// Sets the drag threshold (in pixels).
						threshold: 10

				},

			// If set to a valid selector , prevents key/mouse events from bubbling from these elements.
				excludeSelector: 'input:focus, select:focus, textarea:focus, audio, video, iframe',

			// Link scroll speed.
				linkScrollSpeed: 1000

		};

	// Skel.
		skel.breakpoints({
			xlarge: '(max-width: 1680px)',
			large: '(max-width: 1280px)',
			medium: '(max-width: 980px)',
			small: '(max-width: 736px)',
			xsmall: '(max-width: 480px)',
			xxsmall: '(max-width: 360px)',
			short: '(min-aspect-ratio: 16/7)',
			xshort: '(min-aspect-ratio: 16/6)'
		});

	// Ready event.
		$(function() {

			// Vars.
				var	$window = $(window),
					$document = $(document),
					$body = $('body'),
					$html = $('html'),
					$bodyHtml = $('body,html'),
					$wrapper = $('#wrapper');

			// Disable animations/transitions until the page has loaded.
				$body.addClass('is-loading');

				$window.on('load', function() {
					window.setTimeout(function() {
						$body.removeClass('is-loading');
					}, 100);
				});

			// Tweaks/fixes.

				// Mobile: Revert to native scrolling.
					if (skel.vars.mobile) {

						// Disable all scroll-assist features.
							settings.keyboardShortcuts.enabled = false;
							settings.scrollWheel.enabled = false;
							settings.scrollZones.enabled = false;
							settings.dragging.enabled = false;

						// Re-enable overflow on body.
							$body.css('overflow-x', 'auto');

					}

				// IE: Various fixes.
					if (skel.vars.browser == 'ie') {

						// Enable IE mode.
							$body.addClass('is-ie');

						// Page widths.
							$window
								.on('load resize', function() {

									// Calculate wrapper width.
										var w = 0;

										$wrapper.children().each(function() {
											w += $(this).width();
										});

									// Apply to page.
										$html.css('width', w + 'px');

								});

					}

				// Polyfill: Object fit.
					if (!skel.canUse('object-fit')) {

						$('.image[data-position]').each(function() {

							var $this = $(this),
								$img = $this.children('img');

							// Apply img as background.
								$this
									.css('background-image', 'url("' + $img.attr('src') + '")')
									.css('background-position', $this.data('position'))
									.css('background-size', 'cover')
									.css('background-repeat', 'no-repeat');

							// Hide img.
								$img
									.css('opacity', '0');

						});

					}

			// Keyboard shortcuts.
				if (settings.keyboardShortcuts.enabled)
					(function() {

						$wrapper

							// Prevent keystrokes inside excluded elements from bubbling.
								.on('keypress keyup keydown', settings.excludeSelector, function(event) {

									// Stop propagation.
										event.stopPropagation();

								});

						$window

							// Keypress event.
								.on('keydown', function(event) {

									var scrolled = false;

									switch (event.keyCode) {

										// Left arrow.
											case 37:
												$document.scrollLeft($document.scrollLeft() - settings.keyboardShortcuts.distance);
												scrolled = true;
												break;

										// Right arrow.
											case 39:
												$document.scrollLeft($document.scrollLeft() + settings.keyboardShortcuts.distance);
												scrolled = true;
												break;

										// Page Up.
											case 33:
												$document.scrollLeft($document.scrollLeft() - $window.width() + 100);
												scrolled = true;
												break;

										// Page Down, Space.
											case 34:
											case 32:
												$document.scrollLeft($document.scrollLeft() + $window.width() - 100);
												scrolled = true;
												break;

										// Home.
											case 36:
												$document.scrollLeft(0);
												scrolled = true;
												break;

										// End.
											case 35:
												$document.scrollLeft($document.width());
												scrolled = true;
												break;

									}

									// Scrolled?
										if (scrolled) {

											// Prevent default.
												event.preventDefault();
												event.stopPropagation();

											// Stop link scroll.
												$bodyHtml.stop();

										}

								});

					})();

            // Scroll wheel.
			// WARN: broke scroll so, comment it. 禁止了滚轮事件，于是注释它
            if (false) // (settings.scrollWheel.enabled)
                (function() {

                    // Based on code by @miorel + @pieterv of Facebook (thanks guys :)
                    // github.com/facebook/fixed-data-table/blob/master/src/vendor_upstream/dom/normalizeWheel.js
                    var normalizeWheel = function(event) {

                        var	pixelStep = 10,
                            lineHeight = 40,
                            pageHeight = 800,
                            sX = 0,
                            sY = 0,
                            pX = 0,
                            pY = 0;

                        // Legacy.
                        if ('detail' in event)
                            sY = event.detail;
                        else if ('wheelDelta' in event)
                            sY = event.wheelDelta / -120;
                        else if ('wheelDeltaY' in event)
                            sY = event.wheelDeltaY / -120;

                        if ('wheelDeltaX' in event)
                            sX = event.wheelDeltaX / -120;

                        // Side scrolling on FF with DOMMouseScroll.
                        if ('axis' in event
                            &&	event.axis === event.HORIZONTAL_AXIS) {
                            sX = sY;
                            sY = 0;
                        }

                        // Calculate.
                        pX = sX * pixelStep;
                        pY = sY * pixelStep;

                        if ('deltaY' in event)
                            pY = event.deltaY;

                        if ('deltaX' in event)
                            pX = event.deltaX;

                        if ((pX || pY)
                            &&	event.deltaMode) {

                            if (event.deltaMode == 1) {
                                pX *= lineHeight;
                                pY *= lineHeight;
                            }
                            else {
                                pX *= pageHeight;
                                pY *= pageHeight;
                            }

                        }

                        // Fallback if spin cannot be determined.
                        if (pX && !sX)
                            sX = (pX < 1) ? -1 : 1;

                        if (pY && !sY)
                            sY = (pY < 1) ? -1 : 1;

                        // Return.
                        return {
                            spinX: sX,
                            spinY: sY,
                            pixelX: pX,
                            pixelY: pY
                        };

                    };

                    // Wheel event.
                    $body.on('wheel', function(event) {

                        // Disable on <=small.
                        if (skel.breakpoint('small').active)
                            return;

                        // Prevent default.
                        event.preventDefault();
                        event.stopPropagation();

                        // Stop link scroll.
                        $bodyHtml.stop();

                        // Calculate delta, direction.
                        var	n = normalizeWheel(event.originalEvent),
                            x = (n.pixelX != 0 ? n.pixelX : n.pixelY),
                            delta = Math.min(Math.abs(x), 150) * settings.scrollWheel.factor,
                            direction = x > 0 ? 1 : -1;

                        // Scroll page.
                        $document.scrollLeft($document.scrollLeft() + (delta * direction));

                    });

                })();


			// Scroll zones.
				if (settings.scrollZones.enabled)
					(function() {

						var	$left = $('<div class="scrollZone left"></div>'),
							$right = $('<div class="scrollZone right"></div>'),
							$zones = $left.add($right),
							paused = false,
							intervalId = null,
							direction,
							activate = function(d) {

								// Disable on <=small.
									if (skel.breakpoint('small').active)
										return;

								// Paused? Bail.
									if (paused)
										return;

								// Stop link scroll.
									$bodyHtml.stop();

								// Set direction.
									direction = d;

								// Initialize interval.
									clearInterval(intervalId);

									intervalId = setInterval(function() {
										$document.scrollLeft($document.scrollLeft() + (settings.scrollZones.speed * direction));
									}, 25);

							},
							deactivate = function() {

								// Unpause.
									paused = false;

								// Clear interval.
									clearInterval(intervalId);

							};

						$zones
							.appendTo($wrapper)
							.on('mouseleave mousedown', function(event) {
								deactivate();
							});

						$left
							.css('left', '0')
							.on('mouseenter', function(event) {
								activate(-1);
							});

						$right
							.css('right', '0')
							.on('mouseenter', function(event) {
								activate(1);
							});

						$wrapper
							.on('---pauseScrollZone', function(event) {

								// Pause.
									paused = true;

								// Unpause after delay.
									setTimeout(function() {
										paused = false;
									}, 500);

							});

					})();

			// Dragging.
				if (settings.dragging.enabled)
					(function() {

						var dragging = false,
							dragged = false,
							distance = 0,
							startScroll,
							momentumIntervalId, velocityIntervalId,
							startX, currentX, previousX,
							velocity, direction;

						$wrapper

							// Prevent image drag and drop.
								.on('mouseup mousemove mousedown', '.image, img', function(event) {
									event.preventDefault();
								})

							// Prevent mouse events inside excluded elements from bubbling.
								.on('mouseup mousemove mousedown', settings.excludeSelector, function(event) {

									// Prevent event from bubbling.
										event.stopPropagation();

									// End drag.
										dragging = false;
										$wrapper.removeClass('is-dragging');
										clearInterval(velocityIntervalId);
										clearInterval(momentumIntervalId);

									// Pause scroll zone.
										$wrapper.triggerHandler('---pauseScrollZone');

								})

							// Mousedown event.
								.on('mousedown', function(event) {

									// Disable on <=small.
										if (skel.breakpoint('small').active)
											return;

									// Clear momentum interval.
										clearInterval(momentumIntervalId);

									// Stop link scroll.
										$bodyHtml.stop();

									// Start drag.
										dragging = true;
										$wrapper.addClass('is-dragging');

									// Initialize and reset vars.
										startScroll = $document.scrollLeft();
										startX = event.clientX;
										previousX = startX;
										currentX = startX;
										distance = 0;
										direction = 0;

									// Initialize velocity interval.
										clearInterval(velocityIntervalId);

										velocityIntervalId = setInterval(function() {

											// Calculate velocity, direction.
												velocity = Math.abs(currentX - previousX);
												direction = (currentX > previousX ? -1 : 1);

											// Update previous X.
												previousX = currentX;

										}, 50);

								})

							// Mousemove event.
								.on('mousemove', function(event) {

									// Not dragging? Bail.
										if (!dragging)
											return;

									// Velocity.
										currentX = event.clientX;

									// Scroll page.
										$document.scrollLeft(startScroll + (startX - currentX));

									// Update distance.
										distance = Math.abs(startScroll - $document.scrollLeft());

									// Distance exceeds threshold? Disable pointer events on all descendents.
										if (!dragged
										&&	distance > settings.dragging.threshold) {

											$wrapper.addClass('is-dragged');

											dragged = true;

										}

								})

							// Mouseup/mouseleave event.
								.on('mouseup mouseleave', function(event) {

									var m;

									// Not dragging? Bail.
										if (!dragging)
											return;

									// Dragged? Re-enable pointer events on all descendents.
										if (dragged) {

											setTimeout(function() {
												$wrapper.removeClass('is-dragged');
											}, 100);

											dragged = false;

										}

									// Distance exceeds threshold? Prevent default.
										if (distance > settings.dragging.threshold)
											event.preventDefault();

									// End drag.
										dragging = false;
										$wrapper.removeClass('is-dragging');
										clearInterval(velocityIntervalId);
										clearInterval(momentumIntervalId);

									// Pause scroll zone.
										$wrapper.triggerHandler('---pauseScrollZone');

									// Initialize momentum interval.
										if (settings.dragging.momentum > 0) {

											m = velocity;

											momentumIntervalId = setInterval(function() {

												// Scroll page.
													$document.scrollLeft($document.scrollLeft() + (m * direction));

												// Decrease momentum.
													m = m * settings.dragging.momentum;

												// Negligible momentum? Clear interval and end.
													if (Math.abs(m) < 1)
														clearInterval(momentumIntervalId);

											}, 15);

										}

								});

					})();

			// Link scroll.
				$wrapper
					.on('mousedown mouseup', 'a[href^="#"]', function(event) {

						// Stop propagation.
							event.stopPropagation();

					})
					.on('click', 'a[href^="#"]', function(event) {

						var	$this = $(this),
							href = $this.attr('href'),
							$target, x, y;

						// Get target.
							if (href == '#'
							||	($target = $(href)).length == 0)
								return;

						// Prevent default.
							event.preventDefault();
							event.stopPropagation();

						// Calculate x, y.
							if (skel.breakpoint('small').active) {

								x = $target.offset().top - (Math.max(0, $window.height() - $target.outerHeight()) / 2);
								y = { scrollTop: x };

							}
							else {

								x = $target.offset().left - (Math.max(0, $window.width() - $target.outerWidth()) / 2);
								y = { scrollLeft: x };

							}

						// Scroll.
							$bodyHtml
								.stop()
								.animate(
									y,
									settings.linkScrollSpeed,
									'swing'
								);

					});

			// Gallery.
				$('.gallery')
					.on('click', 'a', function(event) {

						var $a = $(this),
							$gallery = $a.parents('.gallery'),
							$modal = $gallery.children('.modal'),
							$modalImg = $modal.find('img'),
							href = $a.attr('href');

						// Not an image? Bail.
							if (!href.match(/\.(jpg|gif|png|mp4)$/))
								return;

						// Prevent default.
							event.preventDefault();
							event.stopPropagation();

						// Locked? Bail.
							if ($modal[0]._locked)
								return;

						// Lock.
							$modal[0]._locked = true;

						// Set src.
							$modalImg.attr('src', href);

						// Set visible.
							$modal.addClass('visible');

						// Focus.
							$modal.focus();

						// Delay.
							setTimeout(function() {

								// Unlock.
									$modal[0]._locked = false;

							}, 600);

					})
					.on('click', '.modal', function(event) {

						var $modal = $(this),
							$modalImg = $modal.find('img');

						// Locked? Bail.
							if ($modal[0]._locked)
								return;

						// Already hidden? Bail.
							if (!$modal.hasClass('visible'))
								return;

						// Stop propagation.
							event.stopPropagation();

						// Lock.
							$modal[0]._locked = true;

						// Clear visible, loaded.
							$modal
								.removeClass('loaded')

						// Delay.
							setTimeout(function() {

								$modal
									.removeClass('visible')

								// Pause scroll zone.
									$wrapper.triggerHandler('---pauseScrollZone');

								setTimeout(function() {

									// Clear src.
										$modalImg.attr('src', '');

									// Unlock.
										$modal[0]._locked = false;

									// Focus.
										$body.focus();

								}, 475);

							}, 125);

					})
					.on('keypress', '.modal', function(event) {

						var $modal = $(this);

						// Escape? Hide modal.
							if (event.keyCode == 27)
								$modal.trigger('click');

					})
					.on('mouseup mousedown mousemove', '.modal', function(event) {

						// Stop propagation.
							event.stopPropagation();

					})
					.prepend('<div class="modal" tabIndex="-1"><div class="inner"><img src="" /></div></div>')
						.find('img')
							.on('load', function(event) {

								var $modalImg = $(this),
									$modal = $modalImg.parents('.modal');

								setTimeout(function() {

									// No longer visible? Bail.
										if (!$modal.hasClass('visible'))
											return;

									// Set loaded.
										$modal.addClass('loaded');

								}, 275);

							});

		});

})(jQuery);